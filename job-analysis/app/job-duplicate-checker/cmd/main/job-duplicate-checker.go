package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	event "github.com/nadzir/scratchpad-go/job-analysis/pkg/event/jobEvent"
	queue "github.com/nadzir/scratchpad-go/job-analysis/pkg/queue/receive"
)

type jobHashedMap struct {
	jobDescriptionMap map[string]event.JobInfo
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (hashedMap *jobHashedMap) insertToMap(jobInfo event.JobInfo) event.JobInfo {

	desc := jobInfo.Description
	jobHashed := GetMD5Hash(desc)

	value := hashedMap.jobDescriptionMap[jobHashed]
	if value != (event.JobInfo{}) {
		return jobInfo
	} else {
		hashedMap.jobDescriptionMap[jobHashed] = jobInfo
		return event.JobInfo{}
	}
}

func main() {
	newJobChannel := make(chan []byte)
	go queue.StartReceiver(newJobChannel)

	// Init job Hashed map
	jobHashedMap := jobHashedMap{make(map[string]event.JobInfo)}

	for {
		select {
		case msg1 := <-newJobChannel:
			var jobEvent event.JobEvent
			json.Unmarshal(msg1, &jobEvent)
			// fmt.Println(jobEvent.Name, jobEvent.JobID)
			jobInfo := jobEvent.Payload
			// jobDesc := jobInfo.Description
			jobInfoResults := jobHashedMap.insertToMap(jobInfo)
			if jobInfoResults != (event.JobInfo{}) &&
				jobInfoResults.JobURL != jobInfo.JobURL {
				fmt.Println()
				fmt.Println("Duplicated result found")
				fmt.Println("Result 1")
				fmt.Println(jobInfoResults.JobTitle)
				fmt.Println(jobInfoResults.CompanyName)
				fmt.Println(jobInfoResults.JobURL)
				fmt.Println("Result 2")
				fmt.Println(jobInfo.JobTitle)
				fmt.Println(jobInfo.CompanyName)
				fmt.Println(jobInfo.JobURL)
			}
		}
	}
}
