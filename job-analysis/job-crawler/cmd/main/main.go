package main

import (
	"fmt"

	"github.com/nadzir/scratchpad-go/job-analysis/job-crawler/internal/app/crawl/jobcrawler"
)

func main() {
	go jobcrawler.BeginCrawl()

	fmt.Scanln()
}
