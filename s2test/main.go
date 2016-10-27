package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/golang/geo/s2"
)

func main() {
	fmt.Println("This is a s2 test app")

	// init the DB
	SetupSiteBolt()

	//   db, err := bolt.Open("my.db", 0600, nil)
	db, err := bolt.Open("sites.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Compute the CellID for lat, lng
	c := s2.CellIDFromLatLng(s2.LatLngFromDegrees(48.8, 2.0))

	// store the uint64 value of c to its bigendian binary form
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, uint64(c))

	// put the keys in
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URIBucket"))
		err := b.Put(key, []byte("Site 1"))
		return err
	})

	// Compute the CellID for lat, lng
	// c2 := s2.CellIDFromLatLng(s2.LatLngFromDegrees(49.30, 2.7)) // will be in results
	c2 := s2.CellIDFromLatLng(s2.LatLngFromDegrees(49.40, 2.7)) // is not in the results

	// store the uint64 value of c to its bigendian binary form
	key2 := make([]byte, 8)
	binary.BigEndian.PutUint64(key2, uint64(c2))

	// put the keys in
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URIBucket"))
		err := b.Put(key2, []byte("Site 2"))
		return err
	})

	db.Close()

	rect := s2.RectFromLatLng(s2.LatLngFromDegrees(48.99, 1.852))
	rect = rect.AddPoint(s2.LatLngFromDegrees(48.68, 2.75))

	rc := &s2.RegionCoverer{MaxLevel: 20, MaxCells: 8}
	r := s2.Region(rect.CapBound())
	covering := rc.Covering(r)

	for _, c := range covering {
		citiesInCellID(c)
	}

}

func citiesInCellID(c s2.CellID) {
	// compute min & max limits for c
	bmin := make([]byte, 8)
	bmax := make([]byte, 8)
	binary.BigEndian.PutUint64(bmin, uint64(c.RangeMin()))
	binary.BigEndian.PutUint64(bmax, uint64(c.RangeMax()))

	// perform a range lookup in the DB from bmin key to bmax key, cur is our DB cursor
	db, err := bolt.Open("sites.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var cell s2.CellID

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URIBucket"))
		cur := b.Cursor()

		for k, v := cur.Seek(bmin); k != nil && bytes.Compare(k, bmax) <= 0; k, v = cur.Next() {

			fmt.Println("Ready  in loop")

			buf := bytes.NewReader(k)
			binary.Read(buf, binary.BigEndian, &cell)

			// Read back a city
			ll := cell.LatLng()
			lat := float64(ll.Lat.Degrees())
			lng := float64(ll.Lng.Degrees())
			name := string(v)
			fmt.Println(lat, lng, name)
		}

		return nil
	})

}

func SetupSiteBolt() {

	db, err := bolt.Open("sites.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// You can also create a bucket only if it doesn't exist by using the Tx.CreateBucketIfNotExists()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("URIBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		log.Printf("Bucket created %v", b.FillPercent)
		return nil
	})
}
