package teams

import (
	"database/sql"
	"github.com/stevepartridge/baseballga.me-server-go/games"
	"github.com/stevepartridge/go/db"
	"github.com/stevepartridge/go/log"
	"strings"
	"time"
)

const (
	TEAM_SELECT = `
    teams.id,
    teams.name,
    teams.city,
    teams.code,
    teams.domain,
    teams.timezone,
    teams.timezone_offset,
    teams.league,
    teams.division,
    teams.mlb_id,
    teams.mlb_venue_id,
    teams.mlb_file_code,
    teams.updated_at,
    teams.created_at
  `
)

type Team struct {
	Id             int          `json:"id"`
	Name           string       `json:"name"`
	City           string       `json:"city"`
	Code           string       `json:"code"`
	Domain         string       `json:"domain"`
	Timezone       string       `json:"timezone"`
	TimezoneOffset int64        `json:"timezone_offset"`
	League         string       `json:"league"`
	Division       string       `json:"division"`
	Games          []games.Game `json:"games"`
	MlbId          string       `json:"mlb_id"`
	MlbVenueId     string       `json:"mlb_venue_id"`
	MlbFileCode    string       `json:"mlb_file_code"`
	UpdatedAt      time.Time    `json:"updated_at"`
	CreatedAt      time.Time    `json:"created_at"`
}

type TeamResult struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	City           sql.NullString `json:"city"`
	Code           sql.NullString `json:"code"`
	Domain         sql.NullString `json:"domain"`
	Timezone       sql.NullString `json:"timezone"`
	TimezoneOffset sql.NullInt64  `json:"timezone_offset"`
	League         sql.NullString `json:"league"`
	Division       sql.NullString `json:"division"`
	MlbId          sql.NullString `json:"mlb_id"`
	MlbVenueId     sql.NullString `json:"mlb_venue_id"`
	MlbFileCode    sql.NullString `json:"mlb_file_code"`
	UpdatedAt      interface{}    `json:"updated_at"`
	CreatedAt      interface{}    `json:"created_at"`
}

func parseTeamsFromRows(rows *sql.Rows) []Team {
	_teams := make([]Team, 0)
	if rows == nil {
		return _teams
	}
	var teamResult = TeamResult{}
	for rows.Next() {
		err := rows.Scan(
			&teamResult.Id,
			&teamResult.Name,
			&teamResult.City,
			&teamResult.Code,
			&teamResult.Domain,
			&teamResult.Timezone,
			&teamResult.TimezoneOffset,
			&teamResult.League,
			&teamResult.Division,
			&teamResult.MlbId,
			&teamResult.MlbVenueId,
			&teamResult.MlbFileCode,
			&teamResult.UpdatedAt,
			&teamResult.CreatedAt,
		)
		if err == nil {
			_teams = append(_teams, makeTeamFromResult(teamResult))
		} else {
			log.IfError("ridg.central.teams.parseTeamsFromRows", err)
		}
	}
	return _teams
}

func parseTeamFromRows(rows *sql.Rows) Team {
	_teams := parseTeamsFromRows(rows)
	if len(_teams) == 0 {
		return Team{}
	} else {
		return _teams[0]
	}
}

func makeTeamFromResult(result TeamResult) Team {

	team := Team{}
	team.Id = result.Id
	team.Name = result.Name
	team.City = db.Utils.NullStringToString(result.City)
	team.Code = db.Utils.NullStringToString(result.Code)
	team.Domain = db.Utils.NullStringToString(result.Domain)
	team.Timezone = strings.TrimRight(db.Utils.NullStringToString(result.Timezone), " ")
	team.TimezoneOffset = db.Utils.NullInt64ToInt64(result.TimezoneOffset)
	team.League = strings.TrimRight(db.Utils.NullStringToString(result.League), " ")
	team.Division = db.Utils.NullStringToString(result.Division)
	team.Games = []games.Game{}
	team.MlbId = db.Utils.NullStringToString(result.MlbId)
	team.MlbVenueId = db.Utils.NullStringToString(result.MlbVenueId)
	team.MlbFileCode = strings.TrimRight(db.Utils.NullStringToString(result.MlbFileCode), " ")
	team.CreatedAt = db.Utils.ResultInterfaceToTime(result.CreatedAt)
	team.UpdatedAt = db.Utils.ResultInterfaceToTime(result.UpdatedAt)

	return team
}
