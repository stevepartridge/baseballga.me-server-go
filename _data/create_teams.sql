CREATE TABLE teams (
    id bigint NOT NULL,
    name character varying,
    city character varying,
    code character(3),
    domain character varying,
    timezone character(3),
    timezone_offset integer,
    league character(8),
    division character(1),
    mlb_id character(3),
    mlb_venue_id character varying,
    mlb_file_code character(3),
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL
);

ALTER TABLE public.teams OWNER TO {{DB_USER}};

CREATE SEQUENCE teams_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE public.teams_id_seq OWNER TO {{DB_USER}};

ALTER SEQUENCE teams_id_seq OWNED BY teams.id;

ALTER TABLE ONLY teams ALTER COLUMN id SET DEFAULT nextval('teams_id_seq'::regclass);

