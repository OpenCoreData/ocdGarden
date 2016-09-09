package register

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/emicklei/go-restful"
	"github.com/twinj/uuid"
	// "io"
	"log"
	"strings"
)

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/register").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("/dataset/{facet}/{leg}/{site}/{hole}").Filter(basicAuthenticate).To(authorURI))
	return service
}

// TODO need to deal with optional hole value
func authorURI(req *restful.Request, resp *restful.Response) {
	db, err := bolt.Open("catalog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	URL := fmt.Sprintf("http://opencoredata.org/%v/%v/%v/%v", req.PathParameter("facet"), req.PathParameter("leg"), req.PathParameter("site"), req.PathParameter("hole"))

	// check for existing URI for this URL and return it if there is one

	checkval := "0"
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URIBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			//fmt.Printf("key=%s, value=%s\n", k, v)
			// if string.compare(v, URL)
			compare := strings.Contains(string(v), URL)
			if compare {
				checkval = string(k)
			}
		}
		return nil
	})

	if checkval != "0" {
		log.Println("found a duplicate")
		resp.WriteEntity(checkval)
	} else {

		uuid := uuid.NewV4()
		if err != nil {
			log.Printf("Create time uuid failed: %s", err)
		}
		URI := fmt.Sprintf("http://opencoredata.org/id/dataset/%v", uuid.String())

		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("URIBucket"))
			err := b.Put([]byte(URI), []byte(URL))
			return err
		})

		resp.WriteEntity(URI)
	}
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	encoded := req.Request.Header.Get("Authorization")
	// real code does some decoding
	if len(encoded) == 0 || "Basic YWRtaW46YWRtaW4=" != encoded {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, resp)
}
