# csv-filter-challenge-public
# Instructions
1. Click "Use this template" to create a copy of this repository in your personal github account.  
1. Using technology of your choice, complete assignment listed below.
1. Update the README in your new repo with:
    * a `How-To` section containing any instructions needed to execute your program.
    * an `Assumptions` section containing documentation on any assumptions made while interpreting the requirements.
1. Send an email to Scoir (andrew@scoir.com) with a link to your newly created repo containing the completed exercise.

## Expectations
1. This exercise is meant to drive a conversation. 
1. Please invest only enough time needed to demonstrate your approach to problem solving, code design, etc.
1. Within reason, treat your solution as if it would become a production system.

## Assignment
Create a command line application that parses a CSV file and filters the data per user input.

The CSV will contain three fields: `first_name`, `last_name`, and `dob`. The `dob` field will be a date in YYYYMMDD format.

The user should be prompted to filter by `first_name`, `last_name`, or birth year. The application should then accept a name or year and return all records that match the value for the provided filter. 

Example input:
```
first_name,last_name,dob
Bobby,Tables,19700101
Ken,Thompson,19430204
Rob,Pike,19560101
Robert,Griesemer,19640609
```

## How-To

1. Download a Go compiler [here](https://go.dev/dl/) (version 1.18 or greater).
2. `cd` into the root project directory.
3. Run `go build cmd/csvfilter/csvfilter.go`. A binary file (`csvfilter`) should be created.
4. Run `./csvfilter` and follow the prompts :)

NOTE: A `csvfilter` binary compiled for darwin/amd64 is included which will work on most Intel based Macs.

## Assumptions

- Input CSV file conforms to [RFC 4180](https://www.rfc-editor.org/rfc/rfc4180.html) (as required by the Go `csv` package).
    - The first line is always a header and subsequent lines will have the same number of fields as the header.
    - If any record is missing one of `first_name`, `last_name`, or `dob`, the data is malformed and an error should be returned.
- Any errors in data format or integrity should be fatal.
- The `dob` field may not be blank for any record.