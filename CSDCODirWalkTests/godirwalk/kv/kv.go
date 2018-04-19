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
	fmt.Printf("For doc %s I am recording a new event %s \n", file, eventID)
	// fmt.Printf("%s  %s  %s  %s \n", valid, projname, file, measurement)

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

// TODO..  get this the prov vetted by eventIDer
// GetProvCuratedGraph

// TODO..  get the graph that is the roll up of all
// prov sent in.  Which means I need to follow up on the
// URI-Lists and build a graph from them.
// GetProvCommunityGraph

// //GetProvDetails Return the entry for a specific prov record
// func GetProvDetails(eventID string) (string, string, error) {
// 	fmt.Printf("Request content eventID %s \n", eventID)
// 	db := getKVStoreRO()

// 	var provEntry string
// 	var contentType string
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("ProvBucket"))
// 		b2 := tx.Bucket([]byte("LogBucket"))

// 		v := b.Get([]byte(eventID))
// 		provEntry = string(v)

// 		v2 := b2.Get([]byte(eventID))
// 		logLine := strings.Split(string(v2), ",") // 3rd entry is content-type, see NewProvEvent write to this bucket
// 		if len(logLine) == 3 {
// 			contentType = logLine[2]
// 		} else {
// 			contentType = "text/plain" // ??
// 		}

// 		log.Println(logLine)

// 		return nil
// 	})

// 	if err != nil {
// 		log.Println("Error reading from Buckets")
// 	}

// 	err = db.Close()
// 	if err != nil {
// 		log.Println("Error closing database index.db")
// 		log.Println(err)
// 	}

// 	return provEntry, contentType, nil
// }

// // GetProvLog gets all the logged events for a given docID
// func GetProvLog(docID string) (map[string]string, error) {
// 	db := getKVStoreRO()

// 	eventmap := make(map[string]string)

// 	// Logic needed
// 	// 1) loop over IDLinkBucket to find all eventID that match a value of docID
// 	// 2) for each eventID, pull event (value) from LogBucket
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("IDLinkBucket"))
// 		b2 := tx.Bucket([]byte("LogBucket"))
// 		c := b.Cursor()
// 		for k, v := c.First(); k != nil; k, v = c.Next() {
// 			if strings.Contains(string(v), docID) {
// 				v2 := b2.Get(k)
// 				eventmap[string(k)] = string(v2)
// 			}
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		log.Println("Error reading file info from the KV store index.db")
// 		log.Println(err)
// 	}

// 	err = db.Close()
// 	if err != nil {
// 		log.Println("Error closing database index.db")
// 		log.Println(err)
// 	}

// 	return eventmap, err
// }

// // GetDocIDs get all the files in our holding
// func GetDocIDs() []string {
// 	db := getKVStoreRO()

// 	var IDs []string
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("MetaDataBucket"))
// 		c := b.Cursor()
// 		for k, _ := c.First(); k != nil; k, _ = c.Next() {
// 			// log.Printf("key=%s, value=%s\n", k, v)
// 			IDs = append(IDs, string(k))
// 		}
// 		return nil
// 	})

// 	// TODO..  add in doing this for external resources too

// 	if err != nil {
// 		log.Println("Error reading file info from the KV store index.db")
// 		log.Println(err)
// 	}

// 	err = db.Close()
// 	if err != nil {
// 		log.Println("Error closing database index.db")
// 		log.Println(err)
// 	}

// 	return IDs
// }

// // GetResMetaData will get the metadata for a dataset
// func GetResMetaData(docID string) (string, error) {
// 	fmt.Printf("I will get the metadata for docID %s \n", docID)
// 	db := getKVStoreRO()

// 	var jsonld string
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("MetaDataBucket"))
// 		v := b.Get([]byte(docID))
// 		jsonld = string(v)
// 		return nil
// 	})

// 	if err != nil {
// 		log.Println("Error reading file info from the KV store index.db")
// 	}

// 	err = db.Close()
// 	if err != nil {
// 		log.Println("Error closing database index.db")
// 		log.Println(err)
// 	}

// 	return jsonld, err
// }

// // GetResData will get the metadata for a dataset
// func GetResData(docID string) (string, error) {
// 	fmt.Printf("I will get the data for docID %s \n", docID)
// 	db := getKVStoreRO()

// 	var datafile string
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("FileBucket"))
// 		v := b.Get([]byte(docID))
// 		datafile = string(v)
// 		return nil
// 	})

// 	if err != nil {
// 		log.Println("Error reading file info from the KV store index.db")
// 	}

// 	err = db.Close()
// 	if err != nil {
// 		log.Println("Error closing database index.db")
// 		log.Println(err)
// 	}

// 	return datafile, err
// }

// // SetResDataByRef take a URI reference and enters that as a
// // resources in the system.
// func SetResDataByRef(ref string) (string, error) {
// 	fmt.Printf("I will set the data for reference %s \n", ref)
// 	db := getKVStoreRW()

// 	docID := uuid.New().String()

// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("RefBucket"))
// 		err := b.Put([]byte(docID), []byte(ref))
// 		return err
// 	})

// 	if err != nil {
// 		log.Println("Error writing reference info from the KV store index.db Filebucket")
// 	}

// 	err = db.Close()
// 	if err != nil {
// 		log.Println("Error closing database index.db")
// 		log.Println(err)
// 	}

// 	return docID, err
// }

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

	// err = db.Update(func(tx *bolt.Tx) error {
	// 	_, err := tx.CreateBucketIfNotExists([]byte("FileBucket"))
	// 	if err != nil {
	// 		return fmt.Errorf("create bucket: %s", err)
	// 	}
	// 	return nil
	// })

	db.Close()

	return err

}
