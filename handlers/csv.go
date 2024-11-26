package handlers

import (
	"encoding/csv"
	"io"
	"strconv"
)

type CSVHandler struct {
	stdin       *csv.Reader
	stdout      *csv.Writer
	stderr      io.Writer
	calculators map[string]Calculator
}

func NewCSVHandler(stdin io.Reader, stdout, stderr io.Writer, calculators map[string]Calculator) *CSVHandler {
	return &CSVHandler{
		stdin:       csv.NewReader(stdin),
		stdout:      csv.NewWriter(stdout),
		stderr:      stderr,
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
			// TODO
		}
		a, _ := strconv.Atoi(record[0])
		calculator := this.calculators[record[1]]
		b, _ := strconv.Atoi(record[2])
		c := calculator.Calculate(a, b)
		_ = this.stdout.Write(append(record, strconv.Itoa(c)))
	}
	this.stdout.Flush()
	return nil
}
