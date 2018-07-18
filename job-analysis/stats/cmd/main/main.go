package main

import (
	"encoding/json"

	event "github.com/nadzir/scratchpad-go/job-analysis/pkg/event/jobEvent"
	queue "github.com/nadzir/scratchpad-go/job-analysis/pkg/queue/receive"
	log "github.com/sirupsen/logrus"
)

// CompanyStats : Company statistic
type CompanyStats struct {
	company    string
	jobSiteMap map[string]int
	totalCount int
}

// CompanyMap :
type CompanyMap struct {
	companyStatsMap map[string]CompanyStats
}

func (m *CompanyMap) insert(company, source string) {
	coyStats := m.companyStatsMap[company]
	if coyStats.company == company {
		jobSiteMap := coyStats.jobSiteMap
		totalCount := coyStats.totalCount

		sourceCount := jobSiteMap[source]
		jobSiteMap[source] = sourceCount + 1

		m.companyStatsMap[company] = CompanyStats{
			company,
			jobSiteMap,
			totalCount + 1,
		}
	} else {
		jobSiteStats := make(map[string]int)
		jobSiteStats[source] = 1
		m.companyStatsMap[company] = CompanyStats{company, jobSiteStats, 1}
	}
	cs := m.companyStatsMap[company]
	cs.Log()
}

func main() {
	newJobChannel := make(chan []byte)
	go queue.StartReceiver(newJobChannel)

	companyMap := CompanyMap{make(map[string]CompanyStats)}

	for {
		select {
		case msg1 := <-newJobChannel:
			var jobEvent event.JobEvent
			json.Unmarshal(msg1, &jobEvent)
			jobInfo := jobEvent.Payload
			companyMap.insert(jobInfo.CompanyName, jobInfo.Source)
		}
	}
}

// Log : CompanyStats loggin
func (cs *CompanyStats) Log() {
	log.WithFields(log.Fields{
		"job site map": cs.jobSiteMap,
		"total count":  cs.totalCount,
	}).Info(cs.company)
}
