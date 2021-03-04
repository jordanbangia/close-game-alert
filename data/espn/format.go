package espn

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jordanbangia/closegamealert/model"
)

func formatData(e EspnScoreboard) ([]model.Game, error) {
	games := []model.Game{}

	for _, event := range e.Events {
		game, err := formatGame(event)
		if err != nil {
			return []model.Game{}, err
		}
		games = append(games, game)
	}

	return games, nil
}

func formatGame(e Event) (model.Game, error) {
	// this only works for NBA and does not work for over time games
	// 4 could go higher in the case of over times
	timeRemaining := int(e.Status.Clock) + (4-e.Status.Period)*12*60

	game := model.Game{
		Name:          e.ShortName,
		Period:        e.Status.Period,
		TimeRemaining: time.Duration(timeRemaining) * time.Second,
	}

	teams := []model.Team{}

	if len(e.Competitions) > 1 {
		logrus.Warn("this should not happen, an event has more than one competition???")
	}

	competition := e.Competitions[0]
	for _, competitor := range competition.Competitors {
		score, err := strconv.Atoi(competitor.Score)
		if err != nil {
			return model.Game{}, fmt.Errorf("failed to parse score as a string: %w", err)
		}

		teams = append(teams, model.Team{
			Name:  competitor.Team.DisplayName,
			Score: score,
		})
	}
	game.Teams = teams
	return game, nil
}
