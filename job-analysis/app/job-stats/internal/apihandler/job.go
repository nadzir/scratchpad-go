package apihandler

import (
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
)

// JobHandler : API routing
// v1/job/count : get the total job count
func JobHandler(router *mux.Router) {
	router.HandleFunc("/v1/job/count", jobCount).Methods("GET")

	router.HandleFunc("/v1/job/popular", popularJob).Methods("GET")
}

func jobCount(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	jobSource := vars.Get("source")
	date := vars.Get("date")

	json := simplejson.New()
	json.Set("job", jobdb.SelectTotalJobCount(jobSource, date))

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func popularJob(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	jobSource := vars.Get("source")
	date := vars.Get("date")

	json := jobdb.SelectPopularJob(jobSource, date)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
