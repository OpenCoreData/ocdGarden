package kv

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

type FileMeta struct {
	Valid       string
	ProjName    string
	File        string
	Measurement string
}

// NewFileEntry function to create the base information
// func NewFileEntry(docID, provFrag, remoteAddress, contentType string) error {
func NewFileEntry(valid, projname, file, measurement string) error {
	eventID := uuid.New().String()
	log.Printf("For doc %s I am recording a new event %s \n", file, eventID)
	// fmt.Printf("%s  %s  %s  %s \n", valid, projname, file, measurement)

	// TODO
	// DO this as a SHA and check to see if the SHA and path are the same...
	// if so, we have been here before...

	fm := FileMeta{Valid: valid, ProjName: projname, File: file, Measurement: measurement}
	fmb, err := json.Marshal(fm)
	if err != nil {
		log.Println("error in json marshaling")
	}

	db := getKVStoreRW()

	// Log the file
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Assessment"))
		err := b.Put([]byte(eventID), fmb)
		return err
	})

	db.Close()
	return err
}

func GetEntries() []FileMeta {
	db := getKVStoreRO()
	var IDs []FileMeta

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Assessment"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			//log.Printf("key=%s, value=%s\n", k, v)
			dat := FileMeta{}
			if err := json.Unmarshal(v, &dat); err != nil {
				panic(err)
			}
			IDs = append(IDs, dat)
		}
		return nil
	})

	if err != nil {
		log.Println("Error reading file info from the KV store index.db")
		log.Println(err)
	}

	err = db.Close()
	if err != nil {
		log.Println("Error closing database index.db")
		log.Println(err)
	}

	return IDs
}

func getKVStoreRW() *bolt.DB {
	db, err := bolt.Open("./kvStores/index.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	return db
}

func getKVStoreRO() *bolt.DB {
	db, err := bolt.Open("./kvStores/index.db", 0666, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	return db
}

// Init the KV store in case we are starting empty and need some buckets made
// Call from the main program at run time...
// report.GenReport("valid", projname, file, "metadata format Dtube Label_")
func InitKV() error {

	db := getKVStoreRW()

	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Assessment"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Close()

	return err

}
