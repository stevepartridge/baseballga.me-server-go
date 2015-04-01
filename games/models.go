package games

import (
	"database/sql"
	"encoding/json"
	"github.com/stevepartridge/go/db"
	"github.com/stevepartridge/go/log"
	"strings"
	"time"
)

const (
	GAME_SELECT = `
			id,
      mlb_game_id,
      gametime_utc,
      venue_id,
      scheduled_innings,
      gameday,

      original_date,

      created_at,
      updated_at,

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
      `
)

type Game struct {
	Id               string    `json:"id"`
	GameId           string    `json:"game_id"`
	GameTimeUTC      time.Time `json:"gametime_utc"`
	VenueId          string    `json:"venue_id"`
	ScheduledInnings string    `json:"scheduled_innings"`
	GameDayId        string    `json:"gameday"`

	OriginalDate string `json:"original_date"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	Status Status `json:"status"`

	Linescore      Linescore `json:"linescore"`
	WinningPitcher Pitcher   `json:"winning_pitcher"`
	LosingPitcher  Pitcher   `json:"losing_pitcher"`

	HomeTeamId            string `json:"home_team_id"`
	HomeCode              string `json:"home_code"`
	HomeFileCode          string `json:"home_file_code"`
	HomeTeamName          string `json:"home_team_name"`
	HomeNameAbbreviated   string `json:"home_name_abbrev"`
	HomeTeamCity          string `json:"home_team_city"`
	HomeWin               string `json:"home_win"`
	HomeLoss              string `json:"home_loss"`
	HomeGamesBack         string `json:"home_games_back"`
	HomeGamesBackWildcard string `json:"home_games_back_wildcard"`
	HomeTime              string `json:"home_time"`
	HomeAmPm              string `json:"home_ampm"`
	HomeTimeZone          string `json:"home_time_zone"`
	HomeTimeZoneOffset    string `json:"time_zone_hm_lg"`
	HomeDivision          string `json:"home_division"`
	HomeLeagueId          string `json:"home_league_id"`
	HomeSportCode         string `json:"home_sport_code"`

	AwayTeamId            string `json:"away_team_id"`
	AwayCode              string `json:"away_code"`
	AwayFileCode          string `json:"away_file_code"`
	AwayTeamName          string `json:"away_team_name"`
	AwayNameAbbreviated   string `json:"away_name_abbrev"`
	AwayTeamCity          string `json:"away_team_city"`
	AwayWin               string `json:"away_win"`
	AwayLoss              string `json:"away_loss"`
	AwayGamesBack         string `json:"away_games_back"`
	AwayGamesBackWildcard string `json:"away_games_back_wildcard"`
	AwayTime              string `json:"away_time"`
	AwayAmPm              string `json:"away_ampm"`
	AwayTimeZone          string `json:"away_time_zone"`
	AwayTimeZoneOffset    string `json:"time_zone_aw_lg"`
	AwayDivision          string `json:"away_division"`
	AwayLeagueId          string `json:"away_league_id"`
	AwaySportCode         string `json:"away_sport_code"`
}

type GameResult struct {
	Id               string      `json:"id"`
	GameId           string      `json:"game_id"`
	GameTimeUTC      interface{} `json:"gametime_utc"`
	VenueId          string      `json:"venue_id"`
	ScheduledInnings string      `json:"scheduled_innings"`
	GameDayId        string      `json:"gameday"`

	OriginalDate interface{}

	UpdatedAt interface{} `json:"updated_at"`
	CreatedAt interface{} `json:"created_at"`

	Status sql.NullString `json:"status"`

	Linescore      sql.NullString `json:"linescore"`
	WinningPitcher sql.NullString `json:"winning_pitcher"`
	LosingPitcher  sql.NullString `json:"losing_pitcher"`

	HomeTeamId            string `json:"home_team_id"`
	HomeCode              string `json:"home_code"`
	HomeFileCode          string `json:"home_file_code"`
	HomeTeamName          string `json:"home_team_name"`
	HomeNameAbbreviated   string `json:"home_name_abbrev"`
	HomeTeamCity          string `json:"home_team_city"`
	HomeWin               string `json:"home_win"`
	HomeLoss              string `json:"home_loss"`
	HomeGamesBack         string `json:"home_games_back"`
	HomeGamesBackWildcard string `json:"home_games_back_wildcard"`
	HomeTime              string `json:"home_time"`
	HomeAmPm              string `json:"home_ampm"`
	HomeTimeZone          string `json:"home_time_zone"`
	HomeTimeZoneOffset    string `json:"time_zone_hm_lg"`
	HomeDivision          string `json:"home_division"`
	HomeLeagueId          string `json:"home_league_id"`
	HomeSportCode         string `json:"home_sport_code"`

	AwayTeamId            string `json:"away_team_id"`
	AwayCode              string `json:"away_code"`
	AwayFileCode          string `json:"away_file_code"`
	AwayTeamName          string `json:"away_team_name"`
	AwayNameAbbreviated   string `json:"away_name_abbrev"`
	AwayTeamCity          string `json:"away_team_city"`
	AwayWin               string `json:"away_win"`
	AwayLoss              string `json:"away_loss"`
	AwayGamesBack         string `json:"away_games_back"`
	AwayGamesBackWildcard string `json:"away_games_back_wildcard"`
	AwayTime              string `json:"away_time"`
	AwayAmPm              string `json:"away_ampm"`
	AwayTimeZone          string `json:"away_time_zone"`
	AwayTimeZoneOffset    string `json:"time_zone_aw_lg"`
	AwayDivision          string `json:"away_division"`
	AwayLeagueId          string `json:"away_league_id"`
	AwaySportCode         string `json:"away_sport_code"`
}

type Status struct {
	Inning      string `json"inning"`
	InningState string `json"inning_state"`
	Strikes     string `json:"s"`
	Balls       string `json:"b"`
	Outs        string `json:"o"`
	TopInning   string `json:"top_inning"`
	Status      string `json:"status"`
	Ind         string `json:"ind"`
	Reason      string `json:"reason"`
}

type Pitcher struct {
	Id                string `json:"id"`
	FirstName         string `json:"first"`
	LastName          string `json:"last"`
	NameDisplayRoster string `json:"name_display_roster"`
	Number            string `json:"number"`
	Wins              string `json:"wins"`
	Losses            string `json:"losses"`
	Era               string `json:"era"`
}

type Linescore struct {
	Runs   LinescoreItem `json:"r"`
	Hits   LinescoreItem `json:"h"`
	Errors LinescoreItem `json:"e"`

	Strikeouts  LinescoreItem `json:"so"`
	Stolenbases LinescoreItem `json:"sb"`
	Homeruns    LinescoreItem `json:"hr"`

	Innings interface{} `json:"inning"`
}

type LinescoreItem struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

func parseGamesFromRows(rows *sql.Rows) []Game {
	_games := make([]Game, 0)
	if rows == nil {
		return _games
	}
	var gameResult = GameResult{}
	for rows.Next() {
		err := rows.Scan(
			&gameResult.Id,
			&gameResult.GameId,
			&gameResult.GameTimeUTC,
			&gameResult.VenueId,
			&gameResult.ScheduledInnings,
			&gameResult.GameDayId,

			&gameResult.OriginalDate,

			&gameResult.CreatedAt,
			&gameResult.UpdatedAt,

			&gameResult.Status,

			&gameResult.Linescore,
			&gameResult.WinningPitcher,
			&gameResult.LosingPitcher,

			&gameResult.HomeTeamId,
			&gameResult.HomeCode,
			&gameResult.HomeFileCode,
			&gameResult.HomeTeamName,
			&gameResult.HomeNameAbbreviated,
			&gameResult.HomeTeamCity,
			&gameResult.HomeWin,
			&gameResult.HomeLoss,
			&gameResult.HomeGamesBack,
			&gameResult.HomeGamesBackWildcard,
			&gameResult.HomeTime,
			&gameResult.HomeAmPm,
			&gameResult.HomeTimeZone,
			&gameResult.HomeTimeZoneOffset,
			&gameResult.HomeDivision,
			&gameResult.HomeLeagueId,
			&gameResult.HomeSportCode,

			&gameResult.AwayTeamId,
			&gameResult.AwayCode,
			&gameResult.AwayFileCode,
			&gameResult.AwayTeamName,
			&gameResult.AwayNameAbbreviated,
			&gameResult.AwayTeamCity,
			&gameResult.AwayWin,
			&gameResult.AwayLoss,
			&gameResult.AwayGamesBack,
			&gameResult.AwayGamesBackWildcard,
			&gameResult.AwayTime,
			&gameResult.AwayAmPm,
			&gameResult.AwayTimeZone,
			&gameResult.AwayTimeZoneOffset,
			&gameResult.AwayDivision,
			&gameResult.AwayLeagueId,
			&gameResult.AwaySportCode,
		)
		if err == nil {
			_games = append(_games, makeGameFromResult(gameResult))
		} else {
			log.IfError(err)
		}
	}
	return _games
}

func parseGameFromRows(rows *sql.Rows) Game {
	_games := parseGamesFromRows(rows)
	if len(_games) == 0 {
		return Game{}
	} else {
		return _games[0]
	}
}

func makeGameFromResult(result GameResult) Game {

	game := Game{}

	game.Id = result.Id
	game.GameId = result.GameId
	game.GameTimeUTC = db.Utils.ResultInterfaceToTime(result.GameTimeUTC)
	game.VenueId = result.VenueId
	game.ScheduledInnings = result.ScheduledInnings
	game.GameDayId = result.GameDayId

	origDate := db.Utils.ResultInterfaceToTime(result.OriginalDate)
	game.OriginalDate = origDate.Format("2006-01-02")

	var err error
	if result.Linescore.String != "" {
		linescore := Linescore{}
		err = json.Unmarshal([]byte(result.Linescore.String), &linescore)
		if err == nil {
			game.Linescore = linescore
		} else {
			log.Warning("Unable to parse linescore", err, result.GameDayId, result.Linescore.String)
		}
	}

	if result.LosingPitcher.String != "" {
		losingPitcher := Pitcher{}
		err = json.Unmarshal([]byte(result.LosingPitcher.String), &losingPitcher)
		log.IfError(err)
		if err == nil {
			game.LosingPitcher = losingPitcher
		}
	}

	if result.WinningPitcher.String != "" {
		winningPitcher := Pitcher{}
		err = json.Unmarshal([]byte(result.WinningPitcher.String), &winningPitcher)
		log.IfError(err)
		if err == nil {
			game.WinningPitcher = winningPitcher
		}
	}

	if result.Status.String != "" {
		status := Status{}
		err = json.Unmarshal([]byte(result.Status.String), &status)
		log.IfError(err)
		if err == nil {
			game.Status = status
		}
	}

	game.HomeTeamId = strings.Trim(result.HomeTeamId, " ")
	game.HomeCode = strings.Trim(result.HomeCode, " ")
	game.HomeFileCode = strings.Trim(result.HomeFileCode, " ")
	game.HomeTeamName = result.HomeTeamName
	game.HomeNameAbbreviated = result.HomeNameAbbreviated
	game.HomeTeamCity = result.HomeTeamCity
	game.HomeWin = result.HomeWin
	game.HomeLoss = result.HomeLoss
	game.HomeGamesBack = result.HomeGamesBack
	game.HomeGamesBackWildcard = result.HomeGamesBackWildcard
	game.HomeTime = result.HomeTime
	game.HomeAmPm = result.HomeAmPm
	game.HomeTimeZone = strings.Trim(result.HomeTimeZone, " ")
	game.HomeTimeZoneOffset = result.HomeTimeZoneOffset
	game.HomeDivision = strings.Trim(result.HomeDivision, " ")
	game.HomeLeagueId = strings.Trim(result.HomeLeagueId, " ")
	game.HomeSportCode = strings.Trim(result.HomeSportCode, " ")

	game.AwayTeamId = strings.Trim(result.AwayTeamId, " ")
	game.AwayCode = strings.Trim(result.AwayCode, " ")
	game.AwayFileCode = strings.Trim(result.AwayFileCode, " ")
	game.AwayTeamName = result.AwayTeamName
	game.AwayNameAbbreviated = result.AwayNameAbbreviated
	game.AwayTeamCity = result.AwayTeamCity
	game.AwayWin = result.AwayWin
	game.AwayLoss = result.AwayLoss
	game.AwayGamesBack = result.AwayGamesBack
	game.AwayGamesBackWildcard = result.AwayGamesBackWildcard
	game.AwayTime = result.AwayTime
	game.AwayAmPm = result.AwayAmPm
	game.AwayTimeZone = strings.Trim(result.AwayTimeZone, " ")
	game.AwayTimeZoneOffset = result.AwayTimeZoneOffset
	game.AwayDivision = strings.Trim(result.AwayDivision, " ")
	game.AwayLeagueId = strings.Trim(result.AwayLeagueId, " ")
	game.AwaySportCode = strings.Trim(result.AwaySportCode, " ")

	game.UpdatedAt = db.Utils.ResultInterfaceToTime(result.UpdatedAt)
	game.CreatedAt = db.Utils.ResultInterfaceToTime(result.CreatedAt)

	return game
}
