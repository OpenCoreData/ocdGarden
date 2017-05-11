package lookup

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/emicklei/go-restful"
	"log"
)

func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/lookup").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	service.Route(service.GET("/byUUID/{ID}").To(byUUID))
	service.Route(service.GET("/all").To(getAll))

	return service
}

//todo return as JSON (JSON-ld)?
func getAll(req *restful.Request, resp *restful.Response) {

	db, err := bolt.Open("catalog.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URIBucket"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			log.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	resp.WriteEntity("get all ID's")
}

func byUUID(req *restful.Request, resp *restful.Response) {

	db, err := bolt.Open("catalog.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var v []byte
	id := req.PathParameter("ID")

	URI := fmt.Sprintf("http://opencoredata.org/id/dataset/%v", id)

	log.Println(URI)

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URIBucket"))
		v = b.Get([]byte(URI))
		return nil
	})

	resp.WriteEntity(string(v))

}
