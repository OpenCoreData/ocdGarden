package main

import (
	"gopkg.in/mgo.v2"
	// "encoding/json"

	"os"
	// "gopkg.in/mgo.v2/bson"
)

func main() {
	// call mongo and lookup the redirection to use...
	session, err := GetMongoCon()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// QAppend(session)
	// IndexSchema(session)

}

func QAppend(session *mgo.Session) {
	// Optional. Switch the session to a monotonic behavior.
	// csvw := session.DB("expedire").C("featureGeoJson")

	// // Find all Documents CSVW
	// var csvwDocs []ocdServices.CSVWMeta // need feature struct
	// err := csvw.Find(nil).All(&csvwDocs)
	// if err != nil {
	// 	fmt.Printf("this is error %v \n", err)
	// }

	// 	change := mgo.Change{
	// 		Update:    bson.M{"$set": bson.M{"field1": "v1"}},
	// 		ReturnNew: true,
	// 	}
	// 	update_post := Post{}
	// 	info, err := DB.C("posts").FindId(id).Apply(change, &update_job) // Apply is correct

}

func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")
	return mgo.Dial(host)
}
