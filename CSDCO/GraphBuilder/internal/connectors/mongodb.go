package connectors

import (
	"gopkg.in/mgo.v2"
	"os"
)

func GetMongoCon() (*mgo.Session, error) {
	host := os.Getenv("MONGO_HOST")

	return mgo.Dial(host)
}
