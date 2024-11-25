package main

import (
	"log"
	"os"

	"github.com/mdw-smarty/calc-apps2/handlers"
	"github.com/mdw-smarty/calc-lib2"
)

func main() {
	calculator := &calc.Addition{}
	handler := handlers.NewCLIHandler(calculator, os.Stdout)
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
