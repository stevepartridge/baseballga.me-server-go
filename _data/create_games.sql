SET TimeZone = 'UTC';
CREATE TABLE games (
    id bigint NOT NULL,
    mlb_game_id character varying,
    gametime_utc timestamp without time zone,
    original_date date,
    day character varying,
    venue_id character varying,
    scheduled_innings character varying,
    gameday character varying,
    home_id bigint,
    home_mlb_id character(5),
    home_code character(4),
    home_file_code character(5),
    home_team_name character varying,
    home_name_abbrev character varying,
    home_team_city character varying,
    home_win character varying,
    home_loss character varying,
    home_games_back character varying,
    home_games_back_wildcard character varying,
    home_time character varying,
    home_ampm character(2),
    home_time_zone character(5),
    time_zone_hm_lg character varying,
    home_division character(2),
    home_league_id character(4),
    home_sport_code character(5),
    away_id bigint,
    away_mlb_id character(4),
    away_code character(4),
    away_file_code character(5),
    away_team_name character varying,
    away_name_abbrev character varying,
    away_team_city character varying,
    away_win character varying,
    away_loss character varying,
    away_games_back character varying,
    away_games_back_wildcard character varying,
    away_time character varying,
    away_ampm character varying,
    away_time_zone character varying,
    time_zone_aw_lg character varying,
    away_division character varying,
    away_league_id character(4),
    away_sport_code character(5),
    linescore text,
    winning_pitcher text,
    losing_pitcher text,
    status text,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.games OWNER TO {{DB_USER}};

CREATE SEQUENCE games_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.games_id_seq OWNER TO {{DB_USER}};

ALTER SEQUENCE games_id_seq OWNED BY games.id;

ALTER TABLE ONLY games ALTER COLUMN id SET DEFAULT nextval('games_id_seq'::regclass);

CREATE TRIGGER update_game_updated_at_trigger BEFORE UPDATE ON games FOR EACH ROW EXECUTE PROCEDURE update_updated_at();

