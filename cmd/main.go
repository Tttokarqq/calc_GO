package main

import (
	"github.com/MaksaNeNegr/calc_go/application"
	"github.com/MaksaNeNegr/calc_go/demon"
)

func main() {
	demon.a1()
	app := application.New()
	app.Run()
}