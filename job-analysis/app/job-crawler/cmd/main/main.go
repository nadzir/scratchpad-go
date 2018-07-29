package main

import (
	"fmt"
	"os"

	"github.com/nadzir/scratchpad-go/job-analysis/app/job-crawler/internal/app/crawler"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
)

func main() {
	jobdb.CreateJobTable()

	source := os.Args[1]
	go crawler.Begin(source)

	fmt.Scanln()
}
