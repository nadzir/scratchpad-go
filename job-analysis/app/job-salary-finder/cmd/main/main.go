package main

import (
	"database/sql"
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
)

func main() {
	descriptionSQLChannel := make(chan *sql.Rows)
	go selectAllJobDescription(descriptionSQLChannel)
	go processSalary(descriptionSQLChannel)
	fmt.Scanln()
}

func selectAllJobDescription(descriptionSQLChannel chan<- *sql.Rows) {
	query := "select id, description from jobslisting where description <> ''"
	go db.Select(query, descriptionSQLChannel)
}

func processSalary(descriptionSQLChannel <-chan *sql.Rows) {
	for {
		row := <-descriptionSQLChannel
		var id, description string
		row.Scan(&id, &description)

		r, _ := regexp.Compile(`[\$](\d+(?:[\.\,]\d{1,2})?[kK])`)
		salaries := r.FindAllString(description, -1)
		minSalary := ""
		maxSalary := ""
		for _, salary := range salaries {
			if minSalary == "" {
				minSalary = salary
			}
			if maxSalary == "" {
				maxSalary = salary
			}
			if salary <= minSalary {
				minSalary = salary
			}
			if salary >= maxSalary {
				maxSalary = salary
			}
		}
		minSalaryIn := interpreteSalaryString(minSalary)
		maxSalaryIn := interpreteSalaryString(maxSalary)
		updateSalaryTable(id, minSalaryIn, maxSalaryIn)
	}
}

func interpreteSalaryString(salary string) string {
	if salary == "" {
		return ""
	}
	// Convert string to float
	// eg 2.2 => 2.2
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(salary, -1)
	var parsedSalary float64
	for index, num := range nums {
		numFloat, _ := strconv.ParseFloat(num, 64)
		indexFloat := float64(index)
		parsedSalary = parsedSalary + numFloat*math.Pow(10, -indexFloat)
	}

	// Convert K to 1000
	// Eg 2.2K => 2200
	lastChar := salary[len(salary)-1:]
	if lastChar == "k" || lastChar == "K" {
		parsedSalary = parsedSalary * 1000
	}

	return floattostr(parsedSalary)
}

func updateSalaryTable(id, minSalary, maxSalary string) {
	jobdb.UpdateSalary(id, minSalary, maxSalary)
}

func floattostr(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'g', 6, 64)
}
