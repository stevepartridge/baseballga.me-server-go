package games

import (
	// "encoding/json"
	"github.com/stevepartridge/go/db"
	"github.com/stevepartridge/go/log"
	"strconv"
	"time"
)

func GetById(id int) (Game, error) {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	rows, err := conn.Query(`
    SELECT
      `+GAME_SELECT+`
    FROM
      games
    WHERE
      id = $1
    LIMIT 1`,
		id)
	defer rows.Close()
	log.IfError(err)

	if err != nil {
		return Game{}, err
	}

	return parseGameFromRows(rows), nil
}

func GetByDate(date time.Time, offset, limit string) ([]Game, error) {
	log.Debug("GetByDate", date)

	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "50"
	}

	fromDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	toDate := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())

	from := fromDate.UTC().Format("2006-01-02 15:04:05")
	to := toDate.UTC().Format("2006-01-02 15:04:05")

	log.Debug("from", from)
	log.Debug("to", to)

	query := `
    SELECT
    	` + GAME_SELECT + `
    FROM
      games
    WHERE
      gametime_utc >= $1
        AND
      gametime_utc <= $2
    ORDER BY gametime_utc ASC
    LIMIT $3
    OFFSET $4
    `

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()
	rows, err := conn.Query(query, from, to, limit, offset)
	defer rows.Close()
	log.IfError(err)

	return parseGamesFromRows(rows), err
}

func GetByYear(year int, offset, limit string) ([]Game, error) {
	log.Debug("GetByYear", year)
	if offset == "" {
		offset = "0"
	}

	if limit == "" {
		limit = "50"
	}

	from := strconv.Itoa(year) + "-01-01 00:00:00"
	to := strconv.Itoa(year) + "-12-31 23:59:59"

	query := `
    SELECT
    	` + GAME_SELECT + `
    FROM
      games
    WHERE
      gametime_utc >= $1
        AND
      gametime_utc <= $2
    ORDER BY gametime_utc ASC
    LIMIT $3
    OFFSET $4
    `

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()
	rows, err := conn.Query(query, from, to, limit, offset)
	defer rows.Close()
	log.IfError(err)

	return parseGamesFromRows(rows), err

}

func GetByYearForTeamId(year, teamId int) ([]Game, error) {
	log.Debug("GetByYearForTeamId", year, teamId)

	from := strconv.Itoa(year) + "-01-01 00:00:00"
	to := strconv.Itoa(year) + "-12-31 23:59:59"

	query := `
    SELECT
    	` + GAME_SELECT + `
    FROM
      games
    WHERE
      gametime_utc >= $1
        AND
      gametime_utc <= $2
      	AND
      (
      	home_code = (SELECT code FROM teams WHERE id = $3)
      	OR
      	away_code = (SELECT code FROM teams WHERE id = $4)
      )
    ORDER BY gametime_utc ASC
    `

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()
	rows, err := conn.Query(query, from, to, teamId, teamId)
	defer rows.Close()
	log.IfError(err)

	return parseGamesFromRows(rows), err

}
