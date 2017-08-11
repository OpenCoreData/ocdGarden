package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/blevesearch/bleve"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "opencoredata.org/ocdGarden/BleveIndexer/grpcSearchTest/protobufs"
)

const (
	port = ":50051"
)

// server is used to implement package SearchServer.
type server struct{}

// DoSearch implements search
func (s *server) DoSearch(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	results := searchCall(in.Name, in.Index)
	return &pb.SearchReply{Message: "Results: " + results}, nil
}

func main() {
	log.Println("Opening indexes")

	log.Println("Loading grpc server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearchServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// First test function..   opens each time..  not what we want..
// need to open indexes and maintain state
func searchCall(phrase string, searchIndex string) string {
	log.Printf("Search Term: %s \n", phrase)

	// Open all the index files
	// TODO  really should only open the ones I already know will be in the index alias
	index1, err := openIndex("/Users/dfils/Data/OCDDataVolumes/indexes/abstracts.bleve")
	if err != nil {
		log.Printf("Error with index1 alias: %v", err)
	}
	index2, err := openIndex("/Users/dfils/Data/OCDDataVolumes/indexes/csdco.bleve")
	if err != nil {
		log.Printf("Error with index2 alias: %v", err)
	}
	index3, err := openIndex("/Users/dfils/Data/OCDDataVolumes/indexes/janus.bleve")
	if err != nil {
		log.Printf("Error with index3 alias: %v", err)
	}

	var index bleve.IndexAlias

	if searchIndex == "abstracts" {
		index = bleve.NewIndexAlias(index1)
		log.Println("abstract index only")
	}
	if searchIndex == "csdco" {
		index = bleve.NewIndexAlias(index2)
		log.Println("CSDCO index only")
	}
	if searchIndex == "jrso" {
		index = bleve.NewIndexAlias(index3)
		log.Println("JRSO index only")
	} else {
		index = bleve.NewIndexAlias(index1, index2, index3)
		log.Println("All indexes active")
	}

	// Set up query and search
	// query := bleve.NewMatchQuery(phrase)
	query := bleve.NewQueryStringQuery(phrase)
	search := bleve.NewSearchRequestOptions(query, 10, 0, false) // no explanation
	search.Highlight = bleve.NewHighlight()                      // need Stored and IncludeTermVectors in index
	// search.Highlight = bleve.NewHighlightWithStyle("html") // need Stored and IncludeTermVectors in index

	var jsonResults []byte // will hold our results

	// do search and get results
	searchResults, err := index.Search(search)
	if err != nil {
		log.Printf("Error in search call: %v", err)
	} else {
		hits := searchResults.Hits
		jsonResults, err = json.MarshalIndent(hits, " ", " ")
		if err != nil {
			log.Printf("Error with json marshal call: %v", err)
		}

		// testing print loop
		for k, item := range hits {
			fmt.Printf("\n%d: %s, %f, %s, %v\n", k, item.Index, item.Score, item.ID, item.Fragments)
			for key, frag := range item.Fragments {
				fmt.Printf("%s   %s\n", key, frag)
			}
		}
	}

	return string(jsonResults)
}

func openIndex(indexPath string) (bleve.Index, error) {
	var bleveIdx bleve.Index

	if bleveIdx == nil {
		var err error
		bleveIdx, err = bleve.OpenUsing(indexPath, map[string]interface{}{"read_only": true})
		if err != nil {
			return nil, err
		}
	}

	return bleveIdx, nil
}
