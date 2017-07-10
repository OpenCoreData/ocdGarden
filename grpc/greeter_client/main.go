package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pb "opencoredata.org/ocdGarden/grpc/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// c2 := pb.NewFacilityMetadataClient(conn) // connect to the other defined service

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	// Try again
	r, err = c.SayHelloAgain(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	// yaml to schema.org
	r2, err := c.YAMLToSchema(context.Background(), &pb.YAMLFacilityMeta{Yaml: "YAML version"})
	if err != nil {
		log.Fatalf("could not convert: %v", err)
	}
	log.Printf("JSON-LD: %s", r2.Schemaorg)
}
