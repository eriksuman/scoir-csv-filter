package filter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var dobDateFormat = "20060102" // Format string for parsing dob field

// Record represents a row of data
type Record struct {
	FirstName, LastName string
	DateOfBirth         time.Time
}

func (r Record) String() string {
	return fmt.Sprintf("%s, %s, %s", r.FirstName, r.LastName, r.DateOfBirth.Format(dobDateFormat))
}

func (Record) Header() string {
	return "first_name, last_name, dob"
}

var _ Filterer = (*csvFilterer)(nil)

// A Filterer represents a data set that may be queried by first name, last name or birth year.
type Filterer interface {
	// ByFirstName returns the portion of the data set where FirstName matches name
	ByFirstName(name string) ([]Record, error)
	// ByLastName returns the portion of the data set where LastName matches name
	ByLastName(name string) ([]Record, error)
	// ByBirthYear returns the portion of the data set where birth year matches year
	ByBirthYear(year string) ([]Record, error)
}

// csvFilterer implements Filterer
type csvFilterer struct {
	records []Record
}

func (f *csvFilterer) ByFirstName(name string) ([]Record, error) {
	matches := make([]Record, 0, len(f.records))
	for _, record := range f.records {
		if record.FirstName == name {
			matches = append(matches, record)
		}
	}

	return matches, nil
}

func (f *csvFilterer) ByLastName(name string) ([]Record, error) {
	matches := make([]Record, 0, len(f.records))
	for _, record := range f.records {
		if record.LastName == name {
			matches = append(matches, record)
		}
	}

	return matches, nil
}

func (f *csvFilterer) ByBirthYear(year string) ([]Record, error) {
	if len(year) != 4 {
		return nil, errors.New("input year must be 4 characters long")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return nil, fmt.Errorf("converting year to integer: %v", err)
	}

	matches := make([]Record, 0, len(f.records))
	for _, record := range f.records {
		if record.DateOfBirth.Year() == yearInt {
			matches = append(matches, record)
		}
	}

	return matches, nil
}

// New returns a new Filterer initialized with CSV data from csvFilename. The header line in csvFilename must
// contain first_name, last_name, and dob but may be in any order.
func NewCSV(csvFilename string) (Filterer, error) {
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		return nil, fmt.Errorf("trying to open csv file: %v", err)
	}

	rawRecords, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("parsing file %s as csv: %v", csvFilename, err)
	}

	records, err := buildRecords(rawRecords)
	if err != nil {
		return nil, fmt.Errorf("building records from file %s: %v", csvFilename, err)
	}

	return &csvFilterer{
		records: records,
	}, nil
}

// buildRecords performs further processing on raw CSV data
func buildRecords(rawRecords [][]string) ([]Record, error) {
	if len(rawRecords) <= 1 || len(rawRecords[0]) != 3 {
		return nil, errors.New("record set contains no or malformed data")
	}

	fieldIndexSpec := make(map[string]int)
	for i, field := range rawRecords[0] {
		fieldIndexSpec[strings.ToLower(field)] = i
	}

	records := make([]Record, len(rawRecords[1:]))
	for i, rawRecord := range rawRecords[1:] {

		dobStr := rawRecord[fieldIndexSpec["dob"]]
		if len(dobStr) != 8 {
			return nil, fmt.Errorf("record %d has malformed dob: %s", i, dobStr)
		}

		dobTime, err := time.Parse(dobDateFormat, dobStr)
		if err != nil {
			return nil, fmt.Errorf("parsing dob for record %d: %v", i, err)
		}

		records[i] = Record{
			DateOfBirth: dobTime,
			FirstName:   rawRecord[fieldIndexSpec["first_name"]],
			LastName:    rawRecord[fieldIndexSpec["last_name"]],
		}
	}

	return records, nil
}
