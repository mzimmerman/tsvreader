package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
)

// takes file from stdin and outputs to stdout in the correct time format, one date per line
func main() {
	flag.Parse()
	var columns []int
	for _, arg := range flag.Args() {
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("arg is not an integer - %s", arg)
		}
		if i < 0 {
			log.Fatalf("arg needs to be >= 0 - %s", arg)
		}
		columns = append(columns, i)
	}
	if len(columns) == 0 {
		log.Fatalf("Need to specify which columns to output")
	}
	csvReader := csv.NewReader(bufio.NewReader(os.Stdin))
	csvWriter := csv.NewWriter(bufio.NewWriter(os.Stdout))
	defer csvWriter.Flush()
	csvReader.Comma = '\t'
	for {
		data, err := csvReader.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Error reading standard in - %v", err)
		}
		newData := make([]string, len(columns))
		for x, y := range columns {
			newData[x] = data[y]
		}
		err = csvWriter.Write(newData)
		if err != nil {
			log.Fatalf("Error writing standard out - %v", err)
		}
	}
}
