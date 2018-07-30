package jobstreet

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/worker"
)

// Crawl : Begin crawling url for jobstreet
//         Pattern for indeed url :
//         https://www.jobstreet.com.sg/en/job-search/job-vacancy/%d/?src=16&srcr=&ojs=6 where d is 1,2,3..
func Crawl() {
	const numOfPages = 99999
	const numOfWorkers = 10

	urlChannel := make(chan string)

	for w := 0; w <= numOfWorkers; w++ {
		go worker.Crawler(urlChannel, crawlURL)
	}
	// crawlURL("https://www.jobstreet.com.sg/en/job-search/job-vacancy.php?ojs=6")
	for i := 1; i < numOfPages; i++ {
		url := fmt.Sprintf("https://www.jobstreet.com.sg/en/job-search/job-vacancy/%d/?src=16&srcr=&ojs=6", i)
		urlChannel <- url
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

		jobInfo := jobdb.JobInfo{
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
		jobdb.InsertJobTable(jobInfo)
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
