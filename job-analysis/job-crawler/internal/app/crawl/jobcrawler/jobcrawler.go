package jobcrawler

import "github.com/nadzir/scratchpad-go/job-analysis/job-crawler/internal/app/crawl/jobstreet"

func BeginCrawl() {
	jobstreet.Crawl()
}
