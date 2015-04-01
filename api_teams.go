package main

import (
	// "./teams"
	"github.com/stevepartridge/baseballga.me-server-go/games"
	"github.com/stevepartridge/baseballga.me-server-go/teams"
	"github.com/stevepartridge/go/log"
	"github.com/zenazn/goji/web"
	"net/http"
	"strconv"
	"time"
)

func GetTeams(c web.C, res http.ResponseWriter, req *http.Request) {

	offset := c.URLParams["offset"]
	limit := c.URLParams["limit"]

	result, err := teams.Get(offset, limit)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"teams": result, "total": len(result)})
}

func GetGamesByTeamId(c web.C, res http.ResponseWriter, req *http.Request) {

	teamId := c.URLParams["teamId"]

	id, err := strconv.Atoi(teamId)
	log.IfError(err)
	result, err := games.GetByYearForTeamId(time.Now().Year(), id)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"games": result, "total": len(result)})
}

func GetTeamById(c web.C, res http.ResponseWriter, req *http.Request) {

	teamId := c.URLParams["teamId"]

	id, err := strconv.Atoi(teamId)

	result, err := teams.GetById(id)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"team": result})
}

func GetTeamByMlbId(c web.C, res http.ResponseWriter, req *http.Request) {

	mlbId := c.URLParams["mlbId"]

	timezone := req.URL.Query().Get("tz")
	if timezone == "" {
		timezone = "-0400"
	}

	result, err := teams.GetByMlbId(mlbId)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"team": result})
}

func GetTeamByDomain(c web.C, res http.ResponseWriter, req *http.Request) {

	domain := c.URLParams["domain"]

	timezone := req.URL.Query().Get("tz")
	if timezone == "" {
		timezone = "-0400"
	}

	result, err := teams.GetByDomain(domain)
	if err != nil {
		api.ResponseErrorJSON(res, req, 500, []string{err.Error()})
		return
	}
	api.ResponseJSON(res, req, 200, map[string]interface{}{"team": result})
}
