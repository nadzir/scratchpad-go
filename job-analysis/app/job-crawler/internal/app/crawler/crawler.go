package crawler

import (
	"github.com/nadzir/scratchpad-go/job-analysis/app/job-crawler/internal/app/crawler/indeed"
	"github.com/nadzir/scratchpad-go/job-analysis/app/job-crawler/internal/app/crawler/jobstreet"
)

// Begin : Begin crawling
func Begin() {
	go jobstreet.Crawl()
	go indeed.Crawl()
}
