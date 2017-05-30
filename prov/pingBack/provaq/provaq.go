package provaq

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

// New is a function to expose provaq test endpoint
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/provaq").
		Doc("BETA: Not versioned, not for general use.").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("/alpha1/{foo}").To(alpha1).
		Doc("Alpha1:  provaq").
		Param(service.PathParameter("foo", "Foo is foo").DataType("string")).
		Operation("alpha1"))

	service.Route(service.POST("/alpha2").Consumes("text/uri-list").To(alpha2).
		Doc("Alpha2:  provaq").
		Operation("alpha2"))

	return service
}

func alpha2(request *restful.Request, response *restful.Response) {
	// foo := request.PathParameter("foo")

	body, err := ioutil.ReadAll(request.Request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(response, "can't read body", http.StatusBadRequest)
	}

	// do something with the POST data
	// likley convert to triples and write to some end point...
	fmt.Println(string(body))

	// https://golang.org/pkg/net/http/#pkg-constants for my codes....
	response.WriteHeader(http.StatusNoContent)
	response.WriteEntity(string(body)) // remove this....  we are 204
}

func alpha1(request *restful.Request, response *restful.Response) {
	// foo := request.PathParameter("foo")
	data := "A pingback-URI may respond to other requests, but no requirements are imposed on how it responds."
	response.WriteEntity(data)
}
