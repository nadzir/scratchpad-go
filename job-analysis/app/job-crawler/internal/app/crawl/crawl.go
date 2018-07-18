package crawl

import (
	"github.com/nadzir/scratchpad-go/job-analysis/app/job-crawler/internal/app/crawl/indeed"
	"github.com/nadzir/scratchpad-go/job-analysis/app/job-crawler/internal/app/crawl/jobstreet"
)

// Begin : Begin crawling
func Begin() {
	go jobstreet.Crawl()
	go indeed.Crawl()
}
