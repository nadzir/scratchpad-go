package event

import (
	"encoding/json"

	queue "github.com/nadzir/scratchpad-go/job-analysis/pkg/queue/send"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/stringutil"
	log "github.com/sirupsen/logrus"
)

const (
	// CRAWLED_JOB : When a job is crawled
	crawledJob = "crawledJob"
)

// JobInfo : structure for job information
type JobInfo struct {
	Source      string
	CrawledURL  string
	JobURL      string
	JobTitle    string
	CompanyName string
	Description string
	PostingDate string
	ClosingDate string
}

// JobEvent : structure for job event
type JobEvent struct {
	Name    string
	JobID   string
	Payload JobInfo
}

// Log : JobInfo loggin
func (j *JobInfo) Log() {
	log.WithFields(log.Fields{
		"source":      j.Source,
		"crawled url": j.CrawledURL,
		"job url":     j.JobURL,
		"job title":   j.JobTitle,
		"company":     j.CompanyName,
		// "desc":         j.Description,
		"posting date": j.PostingDate,
		"closing date": j.ClosingDate,
	}).Info("Job Info")
}

// CrawledJob : crawl
func CrawledJob(jobInfo JobInfo) {

	jobTitle := jobInfo.JobTitle
	companyName := jobInfo.CompanyName
	jobDesc := jobInfo.Description

	if jobTitle == "" || companyName == "" || jobDesc == "" {
		return
	}
	// Hash job title and companyName
	jobID := stringutil.GetMD5Hash(jobTitle + companyName)

	jobEvent := JobEvent{
		crawledJob,
		jobID,
		jobInfo,
	}

	msg, err := json.Marshal(jobEvent)
	if err != nil {
		panic(err)
	}
	queue.Send(msg)
}
