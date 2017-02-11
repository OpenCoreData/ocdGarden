package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	minio "github.com/minio/minio-go"
	"opencoredata.org/ocdServices/utilities"
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

	// Add logging filter
	wsContainer.Filter(utilities.NCSACommonLogFormatLogger())

	// Add the services
	wsContainer.Add(New())

	// Swagger
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		ApiPath:        "/apidocs.json",
		WebServicesUrl: "http://opencoredata.org"} // localhost:6789
	// SwaggerPath:     "/apidocs/"
	// SwaggerFilePath: "/Users/dfils/src/go/src/opencoredata.org/ocdWeb/static/swagger-ui"}

	// swagger.InstallSwaggerService(config)  // what is this, seen it in use some places
	swagger.RegisterSwaggerService(config, wsContainer)

	// Print out the ports in use and start the services
	log.Printf("Services on localhost:6789")
	// log.Printf("Serving graphql and HTML on localhost:7890/graphql")

	server := &http.Server{Addr: ":7890", Handler: wsContainer}
	server.ListenAndServe()

}

// New is a call to the s3client system
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/api/v1/docrequest").
		Doc("Access Open Core Data Documents").
		Consumes("application/x-www-form-urlencoded").
		Produces("text/plain")

	service.Route(service.GET("/test").To(rockEval).
		Doc("Rock evaluation").
		Param(service.QueryParameter("leg", "Leg of expedition")).
		Param(service.QueryParameter("site", "Site of expedition")).
		Param(service.QueryParameter("hole", "Hole of expedition")).
		Param(service.QueryParameter("core", "Core")).
		Param(service.QueryParameter("section", "Core section")).
		Param(service.QueryParameter("depthtop", "Depth top")).
		Param(service.QueryParameter("depthbottom", "Depth bottom")).
		Operation("Rock evaluation query"))

	return service
}

func rockEval(request *restful.Request, response *restful.Response) {

	endpoint := "172.20.42.161:9000"
	accessKeyID := "E48OPWF0ICVNMF4E3UHN"
	secretAccessKey := "qJdWwMTN4ZyO/jwrmueQRUODM51C+WLzW3efr88U"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Set a bucked called coreimages.
	bucketName := "coreimages"
	// location := "us-east-1"
	objectName := "coreref2.png"

	object, err := minioClient.GetObject(bucketName, objectName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)

	stat, err := object.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := io.CopyN(foo, object, stat.Size); err != nil {
		log.Fatalln(err)
	}

	response.Write(b.Bytes())
}
