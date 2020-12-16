package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/internal/connectors"
	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/internal/query"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	log.Println("Load SQLite3 features and resources into DOC")

	v1, err := readConfig("config", map[string]interface{}{
		"sqlfile": "",
		"bucket":  "",
		"minio": map[string]string{
			"address":   "localhost",
			"port":      "9000",
			"accesskey": "",
			"secretkey": "",
		},
	})
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}

	sqlfile := v1.GetString("sqlfile")
	bucket := v1.GetString("bucket")

	db, err := sql.Open("sqlite3", sqlfile)
	if err != nil {
		log.Panic(err)
	}

	mc := connectors.MinioConnection(v1)

	query.Boreholes(db, mc, bucket)
	query.Projects(db, mc, bucket)
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
