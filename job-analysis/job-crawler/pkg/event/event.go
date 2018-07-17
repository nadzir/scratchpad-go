package event

import (
	"encoding/json"

	queue "github.com/nadzir/scratchpad-go/job-analysis/pkg/queue/send"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/stringutil"
)

const (
	NEW_JOB = "NEW_JOB"
)

type JobInfo struct {
	JobTitle    string
	CrawledUrl  string
	JobUrl      string
	CompanyName string
	Description string
}

type JobEvent struct {
	Name    string
	JobID   string
	Payload JobInfo
}

func CreateJobEvent(jobTitle, companyName, crawledUrl, jobUrl, jobDesc string) {

	if jobTitle == "" || companyName == "" || jobDesc == "" {
		return
	}
	// Hash job title and companyName
	jobID := stringutil.GetMD5Hash(jobTitle + companyName)

	jobInfo := JobInfo{
		jobTitle,
		crawledUrl,
		jobUrl,
		companyName,
		jobDesc,
	}

	jobEvent := JobEvent{
		NEW_JOB,
		jobID,
		jobInfo,
	}

	msg, err := json.Marshal(jobEvent)
	if err != nil {
		panic(err)
	}
	queue.Send(msg)
}
