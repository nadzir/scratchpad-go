package worker

import (
	log "github.com/sirupsen/logrus"
)

// Crawler : Crawl the given url and callback
func Crawler(urlChannel <-chan string, fn func(string)) {
	for url := range urlChannel {
		log.Info(url)
		fn(url)
	}
}

// anotherFunction(f func(string) string) {
