package main

import (
	"github.com/zenazn/goji"
)

type Api struct{}

func (a *Api) Serve() {

	a.CORS()
	a.Routes()

	goji.Serve()

}
