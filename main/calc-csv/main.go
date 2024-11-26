package main

import (
	"log"
	"os"

	"github.com/mdw-smarty/calc-apps2/handlers"
	"github.com/mdw-smarty/calc-lib2"
)

func main() {
	handler := handlers.NewCSVHandler(os.Stdin, os.Stdout, os.Stderr, calculators)
	err := handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
