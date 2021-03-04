package main

import (
	"github.com/jordanbangia/closegamealert/data/espn"
	"github.com/sirupsen/logrus"
)

func main() {

	g := espn.NewGetter()

	games, err := g.Get()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Infof("%+v", games)

}
