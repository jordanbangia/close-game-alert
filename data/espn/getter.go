package espn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jordanbangia/closegamealert/model"
)

const address = "http://site.api.espn.com/apis/site/v2/sports/basketball/nba/scoreboard"

type Getter struct {
	client *http.Client
}

func NewGetter() *Getter {
	return &Getter{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (g *Getter) Get() ([]model.Game, error) {
	resp, err := g.client.Get(address)
	if err != nil {
		return nil, fmt.Errorf("failed to get json: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get json, bad status code: %d", resp.StatusCode)
	}

	data := EspnScoreboard{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	return formatData(data)
}
