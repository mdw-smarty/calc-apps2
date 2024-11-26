package main

import (
	"flag"
	"log"
	"os"

	"github.com/mdw-smarty/calc-apps2/handlers"
	"github.com/mdw-smarty/calc-lib2"
)

func main() {
	var operation string
	flag.StringVar(&operation, "op", "+", "The operation to use.")
	flag.Parse()
	calculator, ok := calculators[operation]
	if !ok {
		log.Fatalf("Unknown operation %s.", operation)
	}
	handler := handlers.NewCLIHandler(calculator, os.Stdout)
	err := handler.Handle(flag.Args())
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
