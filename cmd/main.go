package main

import (
	"github.com/MaksaNeNegr/calc_go/application"
	"github.com/MaksaNeNegr/calc_go/demon"
	// "github.com/MaksaNeNegr/calc_go/vars"
	"os"
	// "time"
)

func main() {
	os.Setenv("pTIME_ADDITION_MS", "1000")  		// время выполнения операции сложения в миллисекундах
	os.Setenv("TIME_SUBTRACTION_MS", "1000")  	// время выполнения операции вычитания в миллисекундах
	os.Setenv("TIME_MULTIPLICATIONS_MS", "1000") 				// время выполнения операции умножения в миллисекундах
	os.Setenv("TIME_DIVISIONS_MS", "1000")						// время выполнения операции деления в миллисекунда
	os.Setenv("COMPUTING_POWER", "2") 						// Количество горутин
	
	app := application.New()
	go demon.Demon_func()
	go app.Run()
	select{}
}