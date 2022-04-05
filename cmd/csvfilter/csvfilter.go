package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eriksuman/scoir/filter"
)

func main() {
	filterer, err := filter.NewCSV("example_input.csv")
	if err != nil {
		printErrorAndExit(fmt.Sprintf("Failed to parse CSV file: %v\n", err))
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter filter parameter [first_name, last_name, year]: ")
	column, err := reader.ReadString('\n')
	if err != nil {
		printErrorAndExit(fmt.Sprintf("Failed to read filter column: %v", err))
	}
	column = strings.TrimSpace(column)

	var queryFunc func(string) ([]filter.Record, error)
	switch column {
	case "first_name":
		queryFunc = filterer.ByFirstName
	case "last_name":
		queryFunc = filterer.ByLastName
	case "year":
		queryFunc = filterer.ByBirthYear
	default:
		printErrorAndExit("Unrecognized filter parameter")
	}

	fmt.Printf("Enter a %s value: ", column)
	value, err := reader.ReadString('\n')
	if err != nil {
		printErrorAndExit(fmt.Sprintf("Failed to read filter value: %v", err))
	}
	value = strings.TrimSpace(value)

	result, err := queryFunc(value)
	if err != nil {
		printErrorAndExit(fmt.Sprintf("Getting result set for query: %v", err))
	}

	fmt.Println(filter.Record{}.Header())
	for _, res := range result {
		fmt.Println(res.String())
	}
}

func printErrorAndExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
