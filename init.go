package main

import (
	"github.com/stevepartridge/go/db"
	"github.com/stevepartridge/go/db/postgres"
	"github.com/stevepartridge/go/env"
)

func init() {
	initEnv()
	initDb()
	if !check() {
		setup()
	}
}

func initDb() {
	db.Pg.Add(postgres.Database{
		"baseballgame",
		env.Get("BASEBALLGAME_DB_HOST"),
		env.Get("BASEBALLGAME_DB_PORT"),
		env.Get("BASEBALLGAME_DB_NAME"),
		env.Get("BASEBALLGAME_DB_USER"),
		env.Get("BASEBALLGAME_DB_PASS"),
		env.Get("BASEBALLGAME_DB_SSLMODE"),
		nil,
	})
}
