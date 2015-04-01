package main

import (
	"github.com/rs/cors"
	"github.com/zenazn/goji"
)

func (a *Api) CORS() {
	// https://github.com/rs/cors/blob/master/examples/goji/server.go
	c := cors.New(cors.Options{
	// AllowedOrigins: []string{"http://foo.com"},
	})

	goji.Use(c.Handler)
}
