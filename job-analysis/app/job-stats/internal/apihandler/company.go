package apihandler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadzir/scratchpad-go/job-analysis/pkg/db/jobdb"
)

// CompanyHandler : API routing
// v1/job/count : get the total job count
func CompanyHandler(router *mux.Router) {
	router.HandleFunc("/v1/company/popular", popularCompany).Methods("GET")
}

func popularCompany(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	jobSource := vars.Get("source")
	date := vars.Get("date")

	json := jobdb.SelectPopularCompany(jobSource, date)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
