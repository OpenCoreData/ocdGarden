package main

import (
	"flag"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "opencoredata.org/ocdGarden/BleveIndexer/grpcSearchTest/protobufs"
)

const (
	address = "localhost:50051"
)

func main() {
	// Get the command line arg or default to "bear lake" a know CSDCO valid serch term
	searchTermPtr := flag.String("term", "", "a simple search string")
	indexPtr := flag.String("index", "", "restrict to a certain index")
	flag.Parse() // don't forget to parse the flags....
	simpleCall(*searchTermPtr, *indexPtr)

	hackishStress() // just trying to rapid fire test options and monitor the server logs
}

func hackishStress() {
	simpleCall("XRF", "csdco")
	simpleCall("XRF~~", "csdco")
	simpleCall("XRF", "jrso")
	simpleCall("XRF~", "jrso")
	simpleCall("XRF~2", "jrso")
	simpleCall("XRF", "abstracts")
	simpleCall("XRF~2", "abstracts")
	simpleCall("XRF", "")
	simpleCall("XRF~2", "")
	simpleCall("thermal", "csdco")
	simpleCall("thermal~~", "csdco")
	simpleCall("thermal", "jrso")
	simpleCall("thermal~", "jrso")
	simpleCall("thermal~2", "jrso")
	simpleCall("thermal", "abstracts")
	simpleCall("thermal~2", "abstracts")
	simpleCall("thermal", "")
	simpleCall("thermal~2", "")
	simpleCall("lake", "csdco")
	simpleCall("lake~~", "csdco")
	simpleCall("lake", "jrso")
	simpleCall("lake~", "jrso")
	simpleCall("lake~2", "jrso")
	simpleCall("lake", "abstracts")
	simpleCall("lake~2", "abstracts")
	simpleCall("lake", "")
	simpleCall("lake~2", "")
}

func simpleCall(term string, index string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSearchClient(conn)

	r, err := c.DoSearch(context.Background(), &pb.SearchRequest{Name: term, Index: index})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	log.Printf("Search Term %s on index %s \n", term, index)
}
