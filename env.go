package main

import "github.com/stevepartridge/go/env"

var vars = map[string]string{

	"BASEBALLGAME_DB_HOST":    "localhost",
	"BASEBALLGAME_DB_PORT":    "5432",
	"BASEBALLGAME_DB_USER":    "local",
	"BASEBALLGAME_DB_PASS":    "password",
	"BASEBALLGAME_DB_NAME":    "baseballgame_development",
	"BASEBALLGAME_DB_SSLMODE": "disable",
}

func initEnv() {
	for key := range vars {
		env.Add(key, vars[key])
	}
}
