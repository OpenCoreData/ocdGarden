package main

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"opencoredata.org/ocdGarden/PROV/pingBack/provaq"
)

func main() {
	wsContainer := restful.NewContainer()

	// CORS
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// Add the services
	wsContainer.Add(provaq.New())

	// Swagger
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		ApiPath:        "/apidocs.json",
		WebServicesUrl: "http://localhost"} // localhost:16789
	// SwaggerPath:     "/apidocs/"
	// SwaggerFilePath: "/Users/dfils/src/go/src/opencoredata.org/ocdWeb/static/swagger-ui"}

	// swagger.InstallSwaggerService(config)  // what is this, seen it in use some places
	swagger.RegisterSwaggerService(config, wsContainer)

	// Print out the ports in use and start the services
	log.Printf("Services on localhost:16789")

	server := &http.Server{Addr: ":16789", Handler: wsContainer}
	server.ListenAndServe()

}
