package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/golang/geo/s2"
	"github.com/kpawlik/geojson"
)

func main() {
	fmt.Println("This is a s2 test app")

	// init the DB
	SetupSiteBolt()

	// enter some points on the map into the BoltDB
	enterDB(48.8, 2.0, "Site 1")
	enterDB(49.30, 2.7, "Site 2")
	enterDB(19.705627232977267, -155.093994140625, "Hilo Hawaii (in poly and rect)")
	enterDB(21.300570216749353, -157.8680419921875, "Honolulu Hawaii  (should not see so far)")      //POINT(-157.8680419921875 21.300570216749353)
	enterDB(20.514981807048372, -155.9893798828125, "Around hawaii (should not be in poly nor rect") // POINT(-155.9893798828125 20.514981807048372)
	enterDB(20.546329665198517, -156.0552978515625, "Point off Maui (should not be in poly)")        // POINT(-156.0552978515625 20.546329665198517)
	enterDB(20.698436036336485, -156.29837036132812, "Kula Forset Reserve Mau")                      // POINT(-156.29837036132812 20.698436036336485)

	// POINT(-156.2200927734375 20.32498944633163)
	// POINT(-154.720458984375 18.870879505128975)
	rect := s2.RectFromLatLng(s2.LatLngFromDegrees(20.32498944633163, -156.2200927734375))
	rect = rect.AddPoint(s2.LatLngFromDegrees(18.870879505128975, -154.720458984375))

	fmt.Println("----  rectangle search  -----")
	rc := &s2.RegionCoverer{MaxLevel: 30, MaxCells: 300}
	r := s2.Region(rect.CapBound())
	covering := rc.Covering(r)

	for _, c := range covering {
		citiesInCellID(c)
	}

	// Hawaii
	//POLYGON((-155.7366943359375 20.47944647508286,-156.5771484375 19.715969839114035,-155.5718994140625 18.725275098649522,-154.522705078125 19.628036391737734,-155.7366943359375 20.47944647508286))

	// trying to figure out how to build a polygon
	// points need to be counter-clockwise ?
	ll1 := s2.LatLngFromDegrees(20.47944647508286, -155.7366943359375)
	ll2 := s2.LatLngFromDegrees(19.715969839114035, -156.5771484375)
	ll3 := s2.LatLngFromDegrees(18.725275098649522, -155.5718994140625)
	ll4 := s2.LatLngFromDegrees(19.628036391737734, -154.522705078125) // first point is last point

	point1 := s2.PointFromLatLng(ll1)
	point2 := s2.PointFromLatLng(ll2)
	point3 := s2.PointFromLatLng(ll3)
	point4 := s2.PointFromLatLng(ll4)

	points := []s2.Point{}
	points = append(points, point1)
	points = append(points, point2)
	points = append(points, point3)
	points = append(points, point4)

	loop := s2.LoopFromPoints(points)
	loops := []*s2.Loop{}
	loops = append(loops, loop)

	fmt.Println("----  loop search (gets too much) -----")

	defaultCoverer := &s2.RegionCoverer{MaxLevel: 30, MaxCells: 6000}
	rg := s2.Region(loop.CapBound())
	cvr := defaultCoverer.Covering(rg)

	// fmt.Println(poly.CapBound())
	for _, c3 := range cvr {
		citiesInCellID(c3)
	}

	fmt.Println("----  poly search  (doesn't work at all it seems) -----")

	poly := s2.PolygonFromLoops(loops)
	rc2 := &s2.RegionCoverer{MaxLevel: 30, MaxCells: 3000}
	r2 := s2.Region(poly.CapBound())
	covering2 := rc2.Covering(r2)
	// fmt.Println(covering2)
	for _, c2 := range covering2 {
		citiesInCellID(c2)
	}

}

func enterDB(lat, long float64, name string) {
	//   db, err := bolt.Open("my.db", 0600, nil)
	db, err := bolt.Open("sites.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Compute the CellID for lat, lng
	c := s2.CellIDFromLatLng(s2.LatLngFromDegrees(lat, long))

	// store the uint64 value of c to its bigendian binary form
	key := make([]byte, 8)
	binary.BigEndian.PutUint64(key, uint64(c))

	// put the keys in
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URIBucket"))
		err := b.Put(key, []byte(name))
		return err
	})

	db.Close()
}

// ref  http://blog.nobugware.com/post/2016/geo_db_s2_geohash_database/
func citiesInCellID(c s2.CellID) {

	// fmt.Println("Ready  in citiesInCellID")

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

func isClockwisePolygon(p geojson.Coordinates) bool {
	sum := 0.0
	for i, coord := range p[:len(p)-1] {
		next := p[i+1]
		sum += float64((next[0] - coord[0]) * (next[1] + coord[1]))
	}
	if sum == 0 {
		return true
	}
	return sum > 0
}

func reversePolygon(p geojson.Coordinates) {
	for i := len(p)/2 - 1; i >= 0; i-- {
		opp := len(p) - 1 - i
		p[i], p[opp] = p[opp], p[i]
	}
}
