package main

import (
	"fmt"

	"github.com/nadzir/scratchpad-go/job-analysis/app/job-crawler/internal/app/crawler"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
)

func main() {
	jobdb.CreateJobTable()
	go crawler.Begin()

	fmt.Scanln()
}
