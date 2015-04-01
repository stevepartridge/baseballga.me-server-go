package main

import (
	// "./games"
	"github.com/stevepartridge/baseballga.me-server-go/games"
	"github.com/stevepartridge/baseballga.me-server-go/teams"
	"github.com/stevepartridge/go/log"
	"github.com/zenazn/goji/web"
	"net/http"
	"strconv"
	"time"
)

func GetGameByDate(c web.C, res http.ResponseWriter, req *http.Request) {

	year := c.URLParams["year"]
	month := c.URLParams["month"]
	day := c.URLParams["day"]

	timezone := req.URL.Query().Get("tz")
	if timezone == "" {
		timezone = "-0400"
	}

	date, err := time.Parse("2006-01-02 15:04:05 -0700", year+"-"+month+"-"+day+" 00:00:00 "+timezone)
	// date.UTC()
	log.IfError(err)

	result, err := games.GetByDate(date, "", "")
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"games": result, "total": len(result)})
}

func GetGameByTeamId(c web.C, res http.ResponseWriter, req *http.Request) {

	teamId := c.URLParams["teamId"]

	timezone := req.URL.Query().Get("tz")
	if timezone == "" {
		timezone = "-0400"
	}

	// date, err := time.Parse("2006-01-02 15:04:05 -0700", year+"-"+month+"-"+day+" 00:00:00 "+timezone)
	// date.UTC()
	// log.IfError(err)
	id, _ := strconv.Atoi(teamId)
	team, err := teams.GetById(id)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}

	result, err := games.GetFromMlbByTeamMlbId(team.MlbId)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"games": result, "total": len(result)})
}

func GetGameByMlbId(c web.C, res http.ResponseWriter, req *http.Request) {

	mlbId := c.URLParams["mlbId"]

	timezone := req.URL.Query().Get("tz")
	if timezone == "" {
		timezone = "-0400"
	}

	result, err := games.GetFromMlbByMlbId(mlbId)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"game": result})
}

func GetGameById(c web.C, res http.ResponseWriter, req *http.Request) {

	gameId := c.URLParams["gameId"]

	timezone := req.URL.Query().Get("tz")
	if timezone == "" {
		timezone = "-0400"
	}

	id, _ := strconv.Atoi(gameId)
	result, err := games.GetById(id)

	if games.GameNeedsUpdateFromMlb(result) {
		result, err = games.GetById(id)
	}

	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"game": result})
}
