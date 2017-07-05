package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/mzimmerman/multicorecsv"
)

var useCommaIn = flag.Bool("useCommaIn", false, "Specify to use comma instead of tab for delimiter on input")
var useCommaOut = flag.Bool("useCommaOut", false, "Specify to use comma instead of tab for delimiter on output")
var useMultiCore = flag.Bool("multicore", false, "Use multicore - lines will be unordered and input data cannot multiline fields")

// takes file from stdin and outputs to stdout in the correct time format, one date per line
func main() {
	flag.Parse()
	var columns []int
	copyAll := false
	for _, arg := range flag.Args() {
		if arg == "all" {
			copyAll = true
			break
		}
		i, err := strconv.Atoi(arg)
		if err != nil {
			log.Fatalf("arg is not an integer - %s", arg)
		}
		if i < 0 {
			log.Fatalf("arg needs to be >= 0 - %s", arg)
		}
		columns = append(columns, i)
	}
	if len(columns) == 0 && !copyAll {
		log.Fatalf("Need to specify which columns to output")
	}
	csvReader := multicorecsv.NewReader(bufio.NewReader(os.Stdin))
	if !*useCommaIn {
		csvReader.Comma = '\t'
	}
	csvWriter := csv.NewWriter(bufio.NewWriter(os.Stdout))
	defer csvWriter.Flush()
	if !*useCommaOut {
		csvWriter.Comma = '\t'
	}
	var newData []string
	for {
		data, err := csvReader.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Error reading standard in - %v", err)
		}

		if copyAll {
			newData = data
		} else {
			if len(newData) != len(columns) {
				newData = make([]string, len(columns))
			}
			for x, y := range columns {
				newData[x] = data[y]
			}
		}
		err = csvWriter.Write(newData)
		if err != nil {
			log.Fatalf("Error writing standard out - %v", err)
		}
	}
}
