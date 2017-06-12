package provaq

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	mgo "gopkg.in/mgo.v2"

	"github.com/emicklei/go-restful"
)

type PingBack struct {
	URI  string
	Prov string
}

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

func alpha1(request *restful.Request, response *restful.Response) {
	// foo := request.PathParameter("foo")
	data := "A pingback-URI may respond to other requests, but no requirements are imposed on how it responds."
	response.WriteEntity(data)
}

func alpha2(request *restful.Request, response *restful.Response) {

	body, err := ioutil.ReadAll(request.Request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(response, "can't read body", http.StatusBadRequest)
	}

	reader := bytes.NewReader(body)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		writeProvMDB(scanner.Text(), getURLBody(scanner.Text()))
	}

	response.WriteHeader(http.StatusNoContent)
}

func writeProvMDB(uri string, body []byte) {
	session, err := mgo.Dial("opencore.dev:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("prov").C("pingback")

	err = c.Insert(&PingBack{uri, string(body)})
	if err != nil {
		log.Fatal(err)
	}

}

func getURLBody(urlstring string) []byte {
	u, err := url.Parse(urlstring)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Accept", "application/json") // oddly the content-type is ignored for the accept header...
	req.Header.Set("Cache-Control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	// convert the function to also return an err
	// and code up some validators...
	// note to W3C, how do we notify a prov provider the sent bad prov?

	return body
}
