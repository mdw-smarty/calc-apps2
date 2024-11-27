package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type CSVHandler struct {
	stdin       *csv.Reader
	stdout      *csv.Writer
	stderr      *log.Logger
	calculators map[string]Calculator
}

func NewCSVHandler(stdin io.Reader, stdout, stderr io.Writer, calculators map[string]Calculator) *CSVHandler {
	return &CSVHandler{
		stdin:       csv.NewReader(stdin),
		stdout:      csv.NewWriter(stdout),
		stderr:      log.New(stderr, "csv: ", log.LstdFlags),
		calculators: calculators,
	}
}

func (this *CSVHandler) Handle() error {
	this.stdin.FieldsPerRecord = 3
	for {
		record, err := this.stdin.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			this.stderr.Println("Reader err:", err)
			return err
		}
		a, err := strconv.Atoi(record[0])
		if err != nil {
			this.stderr.Printf("Invalid operand: [%s] (err: %v)", record[0], err)
			continue
		}
		calculator, ok := this.calculators[record[1]]
		if !ok {
			this.stderr.Printf("Unsupported operation: [%s]", record[1])
			continue
		}
		b, err := strconv.Atoi(record[2])
		if err != nil {
			this.stderr.Printf("Invalid operand: [%s] (err: %v)", record[2], err)
			continue
		}
		c := calculator.Calculate(a, b)
		err = this.stdout.Write(append(record, strconv.Itoa(c)))
		if err != nil {
			this.stderr.Printf("Write err: %v", err)
			return err
		}
	}
	this.stdout.Flush()
	return this.stdout.Error()
}
