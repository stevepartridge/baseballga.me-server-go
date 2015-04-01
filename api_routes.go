package main

import (
	"github.com/zenazn/goji"
)

func (a *Api) Routes() {

	goji.Get("/games/:year/:month/:day", GetGameByDate)
	// goji.Get("/games/month/:year/:month", GetGameByDate)
	// goji.Get("/games/year/:year", GetGameByDate)
	// goji.Get("/games/:gameId", GetGameById)
	goji.Get("/games/mlbid/:mlbId", GetGameByMlbId)
	goji.Get("/games/:gameId", GetGameById)

	goji.Get("/teams/:teamId/games", GetGamesByTeamId)
	goji.Get("/teams/mlbid/:mlbId", GetTeamById)
	goji.Get("/teams/domain/:domain", GetTeamByDomain)
	goji.Get("/teams/:teamId", GetTeamById)

	goji.Get("/teams", GetTeams)

	goji.Get("/", api.Entry)
	goji.NotFound(api.NotFound)
}
