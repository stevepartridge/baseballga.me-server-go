package teams

import (
	"errors"
	"fmt"
	"github.com/stevepartridge/go/db"
	"github.com/stevepartridge/go/log"
	"time"
)

func Get(offset string, limit string) ([]Team, error) {

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "50"
	}

	query := `
    SELECT
      ` + TEAM_SELECT + `
    FROM
      teams
    ORDER BY
      created_at DESC
    LIMIT $1
    OFFSET $2
    `

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()
	rows, err := conn.Query(query, limit, offset)
	defer rows.Close()
	log.IfError(err)

	return parseTeamsFromRows(rows), err
}

func GetById(id int) (Team, error) {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	rows, err := conn.Query(`
    SELECT
      `+TEAM_SELECT+`
    FROM
      teams
    WHERE
      id = $1
    LIMIT 1`,
		id)
	defer rows.Close()
	log.IfError("baseballgame.GetById", err)

	if err != nil {
		return Team{}, err
	}

	return parseTeamFromRows(rows), nil
}

func GetByMlbId(mlbId string) (Team, error) {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	rows, err := conn.Query(`
    SELECT
      `+TEAM_SELECT+`
    FROM
      teams
    WHERE
      code = $1
        OR
      mlb_id = $2
    LIMIT 1`,
		mlbId, mlbId)
	defer rows.Close()
	log.IfError(err)

	if err != nil {
		return Team{}, err
	}

	return parseTeamFromRows(rows), nil
}

func GetByDomain(domain string) (Team, error) {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	rows, err := conn.Query(`
    SELECT
      `+TEAM_SELECT+`
    FROM
      teams
    WHERE
      domain = $1
    LIMIT 1`,
		domain)
	defer rows.Close()
	log.IfError(err)

	if err != nil {
		return Team{}, err
	}

	return parseTeamFromRows(rows), nil
}

func Create(team Team) (Team, error) {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	log.Info("baseballgame.Create", team)
	var id int
	err := conn.QueryRow(`
    INSERT INTO teams (
      name,
      city,
      code,
      domain,
      timezone,
      timezone_offset,
      league,
      division,
      mlb_id,
      mlb_venue_id,
      mlb_file_code,
      updated_at,
      created_at
       )
    VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
    RETURNING id
    `,
		team.Name,
		team.City,
		team.Code,
		team.Domain,
		team.Timezone,
		team.TimezoneOffset,
		team.League,
		team.Division,
		team.MlbId,
		team.MlbVenueId,
		team.MlbFileCode,
		time.Now(),
		time.Now(),
	).Scan(&id)
	log.IfError(err)

	if err != nil {
		c := Team{}
		c.Id = 0
		return c, err
	} else {
		return GetById(id)
	}
}

func Update(team Team) (Team, error) {

	_team, err := GetById(team.Id)
	if err != nil {
		return Team{}, err
	}
	if _team.Id == 0 {
		return Team{}, errors.New(fmt.Sprintf("Team not found by id: %d", team.Id))
	}

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	result, err := conn.Exec(`
    UPDATE teams SET
      name = $1,
      city = $2,
      code = $3,
      domain = $4,
      timezone = $5,
      timezone_offset = $6,
      league = $7,
      division = $8,
      mlb_id = $9,
      mlb_venue_id = $10,
      mlb_file_code = $11
    WHERE
      id = $12
    `,
		team.Name,
		team.City,
		team.Code,
		team.Domain,
		team.Timezone,
		team.TimezoneOffset,
		team.League,
		team.Division,
		team.MlbId,
		team.MlbVenueId,
		team.MlbFileCode,
		team.Id,
	)

	if err != nil {
		log.IfError(err)
		return Team{}, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		log.Info("rows affected: %d", rowsAffected)
		return GetById(team.Id)
	}
}

func DeleteById(id int) error {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	result, err := conn.Exec(`DELETE FROM teams WHERE id = $1`, id)

	if err != nil {
		log.IfError(err)
		return err
	} else {
		rowsAffected, _ := result.RowsAffected()
		log.Info("rows affected: %d", rowsAffected)
		return nil
	}
}
