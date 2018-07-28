package main

import (
	"fmt"

	"github.com/nadzir/scratchpad-go/job-analysis/app/job-crawler/internal/app/crawl"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
)

func main() {
	jobdb.CreateJobTable()
	go crawl.Begin()

	fmt.Scanln()
}
