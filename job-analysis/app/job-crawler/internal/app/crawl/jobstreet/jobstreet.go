package jobstreet

import (
	"fmt"

	"github.com/gocolly/colly"
	event "github.com/nadzir/scratchpad-go/job-analysis/pkg/event/jobEvent"
	log "github.com/sirupsen/logrus"
)

// Crawl : Begin crawling url for jobstreet
//         Pattern for indeed url :
//         https://www.jobstreet.com.sg/en/job-search/job-vacancy/%d/?src=16&srcr=&ojs=6 where d is 1,2,3..
func Crawl() {
	const numOfPages = 100
	// crawlURL("https://www.jobstreet.com.sg/en/job-search/job-vacancy.php?ojs=6")
	for i := 1; i < numOfPages; i++ {
		url := fmt.Sprintf("https://www.jobstreet.com.sg/en/job-search/job-vacancy/%d/?src=16&srcr=&ojs=6", i)
		log.Info(url)
		crawlURL(url)
	}
}

func crawlURL(url string) {
	c := colly.NewCollector()
	var jobLink string

	c.OnHTML("html", func(e *colly.HTMLElement) {
		jobTitle := e.ChildText("#position_title")
		companyName := e.ChildText("#company_name a")
		jobDescription := e.ChildText("#job_description")
		postingDate := e.ChildText("#posting_date")
		closingDate := e.ChildText("#closing_date")

		jobInfo := event.JobInfo{
			"jobstreet",
			url,
			jobLink,
			jobTitle,
			companyName,
			jobDescription,
			postingDate,
			closingDate,
		}

		jobInfo.Log()
		event.CrawledJob(jobInfo)
	})

	c.OnHTML(".panel-body", func(e *colly.HTMLElement) {
		jobLink = e.ChildAttr(".position-title-link", "href")
		if jobLink != "" {
			e.Request.Visit(jobLink)
		}
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit(url)
}
