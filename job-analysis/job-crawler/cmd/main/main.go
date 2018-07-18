package main

import (
	"fmt"

	"github.com/nadzir/scratchpad-go/job-analysis/job-crawler/internal/app/crawl"
)

func main() {
	go crawl.Begin()
	fmt.Scanln()
}
