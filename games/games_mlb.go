package games

import (
	"encoding/json"
	"github.com/stevepartridge/go/db"
	"github.com/stevepartridge/go/log"
	"github.com/stevepartridge/mlb"
	"strconv"
	"strings"
	"time"
)

func GetFromMlbByDate(date time.Time) ([]mlb.Game, error) {

	gms, err := mlb.GetGamesByDate(date)
	if err != nil {
		log.Error(err)
		return gms, err
	}

	return gms, nil
}

func GetFromMlbByMlbId(mlbId string) (mlb.Game, error) {

	game, err := mlb.GetGameByFileCode(mlbId)
	if err != nil {
		log.Error(err)
		return game, err
	}

	return game, nil
}

func GetFromMlbByTeamMlbId(mlbId string) ([]mlb.Game, error) {
	result := make([]mlb.Game, 0)

	date := time.Now()

	gms, err := mlb.GetGamesByDate(date)
	if err != nil {
		log.Error(err)
		return gms, err
	}

	for _, game := range gms {
		if mlbId == game.HomeTeamId || mlbId == game.AwayTeamId {
			result = append(result, game)
		}
	}

	return result, nil
}

func GameNeedsUpdateFromMlb(game Game) bool {
	now := time.Now().UTC()
	if game.UpdatedAt.Unix() < now.Add(time.Duration(1*time.Minute)).Unix() {
		log.Debug("updated at is longer than a minute ago")

		if game.Status.Status != "Final" {
			if game.UpdatedAt.Unix() < game.GameTimeUTC.Unix() || (game.Status.Status != "Preview") {

				parts := strings.Split(game.OriginalDate, "-")
				year, _ := strconv.Atoi(parts[0])
				month, _ := strconv.Atoi(parts[1])
				day, _ := strconv.Atoi(parts[2])
				loc, _ := time.LoadLocation("Local")
				date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
				gms, err := GetFromMlbByDate(date)
				log.IfError(err)
				updated, errs := UpdateGamesFromMlb(gms)
				log.Info(updated, errs)
				return true
			}
		}
	}

	return false
}

func UpdateGamesFromMlb(gamesList []mlb.Game) (int, []error) {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	stmt, err := conn.Prepare(
		`UPDATE
      games
    SET
      gametime_utc = $1,
      venue_id = $2,
      scheduled_innings = $3,
      status = $4,
      linescore = $5,
      winning_pitcher = $6,
      losing_pitcher = $7
    WHERE
      gameday = $8
    `)
	log.IfError(err)

	_errors := make([]error, 0)
	updated := 0

	for _, game := range gamesList {
		linescore, err := json.Marshal(game.Linescore)
		log.IfError(err)

		winningPitcher, _ := json.Marshal(game.WinningPitcher)
		losingPitcher, _ := json.Marshal(game.LosingPitcher)
		status, _ := json.Marshal(game.Status)

		timeStr := game.OriginalDate + " " + game.Time + game.AmPm + " -0400"
		log.Debug("games", timeStr)
		gametime, te := time.Parse("2006/01/02 3:04PM -0700", timeStr)
		log.IfError(te)

		res, err := stmt.Exec(
			gametime.UTC(),
			game.VenueId,
			game.ScheduledInnings,

			string(status),
			string(linescore),
			string(winningPitcher),
			string(losingPitcher),

			game.GameDayId,
		)

		if err != nil || res == nil {

		} else {
			var ra int64
			ra, err := res.RowsAffected()
			if err != nil {
				_errors = append(_errors, err)
				log.IfError(err)
			} else {
				updated = updated + int(ra)
			}
		}
	}

	defer stmt.Close()

	return updated, _errors
}

func SaveGamesFromMlb(gamesList []mlb.Game) (int, []error) {

	conn := db.Pg.Get("baseballgame")
	defer conn.Close()

	stmt, err := conn.Prepare(
		`INSERT INTO games (
          mlb_game_id,
          gametime_utc,
          original_date,
          day,
          venue_id,
          scheduled_innings,
          gameday,

          status,

          linescore,
          winning_pitcher,
          losing_pitcher,

          home_mlb_id,
          home_code,
          home_file_code,
          home_team_name,
          home_name_abbrev,
          home_team_city,
          home_win,
          home_loss,
          home_games_back,
          home_games_back_wildcard,
          home_time,
          home_ampm,
          home_time_zone,
          time_zone_hm_lg,
          home_division,
          home_league_id,
          home_sport_code,

          away_mlb_id,
          away_code,
          away_file_code,
          away_team_name,
          away_name_abbrev,
          away_team_city,
          away_win,
          away_loss,
          away_games_back,
          away_games_back_wildcard,
          away_time,
          away_ampm,
          away_time_zone,
          time_zone_aw_lg,
          away_division,
          away_league_id,
          away_sport_code
        )
        VALUES (
          $1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
          $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
          $21, $22, $23, $24, $25, $26, $27, $28, $29, $30,
          $31, $32, $33, $34, $35, $36, $37, $38, $39, $40,
          $41, $42, $43, $44, $45
        )`)

	log.IfError(err)

	_errors := make([]error, 0)
	saved := 0

	timezones := map[string]string{"ET": "-0400", "CT": "-0500", "MST": "-0600", "MT": "-0600", "PT": "-0700"}

	for _, game := range gamesList {
		linescore, _ := json.Marshal(game.Linescore)
		winningPitcher, _ := json.Marshal(game.WinningPitcher)
		losingPitcher, _ := json.Marshal(game.LosingPitcher)
		status, _ := json.Marshal(game.Status)

		timeStr := game.OriginalDate + " " + game.Time + game.AmPm + " -0400"

		gametime, te := time.Parse("2006/01/02 3:04PM -0700", timeStr)
		log.IfError(te)

		res, err := stmt.Exec(
			game.GameId,

			gametime.UTC(),

			game.OriginalDate, game.Day, game.VenueId,
			game.ScheduledInnings, game.GameDayId,

			string(status),
			string(linescore),
			string(winningPitcher),
			string(losingPitcher),

			game.HomeTeamId, game.HomeCode, game.HomeFileCode, game.HomeTeamName,
			game.HomeNameAbbreviated, game.HomeTeamCity, game.HomeWin, game.HomeLoss, game.HomeGamesBack,
			game.HomeGamesBackWildcard, game.HomeTime, game.HomeAmPm, game.HomeTimeZone, timezones[game.HomeTimeZone],
			game.HomeDivision, game.HomeLeagueId, game.HomeSportCode,

			game.AwayTeamId, game.AwayCode, game.AwayFileCode, game.AwayTeamName,
			game.AwayNameAbbreviated, game.AwayTeamCity, game.AwayWin, game.AwayLoss, game.AwayGamesBack,
			game.AwayGamesBackWildcard, game.AwayTime, game.AwayAmPm, game.AwayTimeZone, timezones[game.AwayTimeZone],
			game.AwayDivision, game.AwayLeagueId, game.AwaySportCode,
		)

		if err != nil || res == nil {

		} else {
			ra, resultErr := res.RowsAffected()
			if resultErr != nil {
				_errors = append(_errors, resultErr)
				log.IfError(resultErr)
			} else {
				saved = saved + int(ra)
			}
		}
	}

	defer stmt.Close()

	return saved, _errors
}
