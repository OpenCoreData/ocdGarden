package main

import (
	"github.com/emicklei/go-restful"
	//"io"
	"log"
	"net/http"
	"opencoredata.org/ocdCatalog/db"
	"opencoredata.org/ocdCatalog/lookup"
	"opencoredata.org/ocdCatalog/register"
)

func main() {
	// setup bolt if it is not already
	db.SetupBolt()

	wsContainer := restful.NewContainer()

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	wsContainer.Add(register.New())
	wsContainer.Add(lookup.New())

	log.Printf("Listening on localhost:7890")

	server := &http.Server{Addr: ":7890", Handler: wsContainer}
	server.ListenAndServe()
}
