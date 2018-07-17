package jobstreet

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/nadzir/scratchpad-go/job-analysis/job-crawler/pkg/event"
)

func Crawl() {

	const numOfPages = 100

	// crawlURL("https://www.jobstreet.com.sg/en/job-search/job-vacancy.php?ojs=6")
	for i := 1; i < numOfPages; i++ {
		url := fmt.Sprintf("https://www.jobstreet.com.sg/en/job-search/job-vacancy/%d/?src=16&srcr=&ojs=6", i)
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

		event.CreateJobEvent(jobTitle, companyName, url, jobLink, jobDescription)
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

	fmt.Println("Crawling", url)
	c.Visit(url)

}
