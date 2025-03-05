package main

import (
	"github.com/MaksaNeNegr/calc_go/application"
	"github.com/MaksaNeNegr/calc_go/demon"
)

func main() {
	// запуск сервера и 
	app := application.New()
	go demon.Demon_func()
	go app.Run()
	select{}
}