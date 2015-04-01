package main

import (
	"github.com/stevepartridge/baseballga.me-server-go/games"
	"github.com/stevepartridge/go/log"
	"github.com/stevepartridge/mlb"
	"time"
)

func check() bool {
	gms, err := games.GetByYear(time.Now().Year(), "", "")
	log.IfError(err)
	return (len(gms) > 10)
	// log.Debug("count", len(gms), (len(gms) > 10))
	// return true
}

func setup() {

	mlbGames, mlbErrors := loadGamesFromMlbByYear(time.Now().Year())
	log.Info(len(mlbGames), len(mlbErrors))
	if len(mlbGames) >= 1 {
		saved, errors := games.SaveGamesFromMlb(mlbGames)
		log.Info("Saved Games:", saved)
		log.Info("Errors Saving Games:", len(errors))
	}

	return
}

func loadGamesFromMlbByYear(year int) ([]mlb.Game, []error) {
	log.Info("Load Games From MLB by Year", year)
	_games := make([]mlb.Game, 0)
	_errors := make([]error, 0)

	startDate := time.Date(year, 3, 1, 0, 0, 0, 0, time.UTC)
	// endDate := time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, 11, 1, 0, 0, 0, 0, time.UTC)

	total := int((endDate.Unix() - startDate.Unix()) / int64(60*60*24))
	currentDate := startDate

	type MlbResponse struct {
		Key   string
		Games []mlb.Game
		Error error
	}

	respChannel := make(chan *MlbResponse)

	for currentDate.Unix() < endDate.Unix() {
		currentDate = currentDate.Add(time.Hour * 24)
		go func(date time.Time) {
			key := date.Format("2006-01-02")
			gms, err := mlb.GetGamesByDate(date)
			respChannel <- &MlbResponse{key, gms, err}
		}(currentDate)
	}

	count := 0
	for {
		select {
		case resp := <-respChannel:
			if resp.Error != nil {
				log.Warning(resp.Key, "Found error", resp.Error)
				_errors = append(_errors, resp.Error)
			}
			log.Info(resp.Key, "Games Found:", len(resp.Games))
			for _, gm := range resp.Games {
				_games = append(_games, gm)
			}
			count = count + 1
			if total == count {
				log.Info("Total Games found: ", len(_games))
				log.Info("Total Errors: ", len(_errors))
				return _games, _errors
			}
		}
	}

}
