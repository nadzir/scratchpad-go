package indeed

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
	log "github.com/sirupsen/logrus"
)

// Crawl : Begin crawling url for indeed
//         Pattern for indeed url :
//         https://www.indeed.com.sg/jobs?q=&l=Singapore&start=%d where d is 10,20,30..
func Crawl() {

	const numOfPages = 99999

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
		jobTitle := e.ChildText("[data-tn-component=jobHeader] .jobtitle")
		companyName := e.ChildText("[data-tn-component=jobHeader] .company")
		jobDescription := e.ChildText("#job_summary")

		jobInfo := jobdb.JobInfo{
			"indeed",
			url,
			jobLink,
			jobTitle,
			companyName,
			jobDescription,
			"",
			"",
		}
		// if jobTitle != "" && companyName != "" && jobDescription != "" {
		jobInfo.Log()
		jobdb.InsertJobTable(jobInfo)
		// }
	})

	c.OnHTML(".row", func(e *colly.HTMLElement) {
		jobLink = e.ChildAttr(".turnstileLink", "href")
		if jobLink != "" {
			e.Request.Visit(jobLink)
		}
	})

	c.Visit(url)
}
