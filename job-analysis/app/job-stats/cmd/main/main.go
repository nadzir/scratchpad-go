package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadzir/scratchpad-go/job-analysis/app/job-stats/internal/apihandler"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	apihandler.JobHandler(router)
	apihandler.CompanyHandler(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	port := ":3100"
	fmt.Println("Listening on port ", port)
	err := http.ListenAndServe(":3100", handler)
	if err != nil {
		log.Fatal(err)
	}
}
