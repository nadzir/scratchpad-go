package indeed

import (
	"fmt"

	"github.com/gocolly/colly"
	event "github.com/nadzir/scratchpad-go/job-analysis/pkg/event/jobEvent"
	log "github.com/sirupsen/logrus"
)

// Crawl : Begin crawling url for indeed
//         Pattern for indeed url :
//         https://www.indeed.com.sg/jobs?q=&l=Singapore&start=%d where d is 10,20,30..
func Crawl() {

	const numOfPages = 100

	// crawlURL("https://www.jobstreet.com.sg/en/job-search/job-vacancy.php?ojs=6")
	for i := 1; i < numOfPages; i++ {
		url := fmt.Sprintf("https://www.indeed.com.sg/jobs?q=&l=Singapore&start=%d", i*10)
		log.Info(url)
		crawlURL(url)
	}
}

func crawlURL(url string) {
	c := colly.NewCollector()
	var jobLink string

	c.OnHTML("html", func(e *colly.HTMLElement) {
		jobTitle := e.ChildText(".jobtitle")
		companyName := e.ChildText(".company")
		jobDescription := e.ChildText("#job_summary")

		jobInfo := event.JobInfo{
			"indeed",
			url,
			jobLink,
			jobTitle,
			companyName,
			jobDescription,
			"",
			"",
		}
		jobInfo.Log()
		event.CrawledJob(jobInfo)
	})

	c.OnHTML(".row", func(e *colly.HTMLElement) {
		jobLink = e.ChildAttr(".jobtitle", "href")
		if jobLink != "" {
			e.Request.Visit(jobLink)
		}
	})

	c.Visit(url)
}
