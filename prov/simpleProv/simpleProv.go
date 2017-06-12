package main

import (
	"fmt"
	"log"
	"os"

	"github.com/knakk/rdf"
)

func main() {
	fmt.Println("simple prov builder")

	tr := []rdf.Triple{}

	// RDF = rdf.Predicate.
	// PROV = Namespace('http://www.w3.org/ns/prov#')
	//  g.add((this_sample, RDF.type, PROV.Entity))
	// 	g.add((ga, RDF.type, PROV.Org))
	// 	qualified_attribution = BNode()
	// 	g.add((qualified_attribution, RDF.type, PROV.Attribution))
	// 	g.add((qualified_attribution, PROV.agent, ga))
	// 	g.add((qualified_attribution, PROV.hadRole, AUROLE.Publisher))
	// 	g.add((this_sample, PROV.qualifiedAttribution, qualified_attribution))

	// # just for visjs
	// g.add((ga, RDF.type, PROV.Agent))
	// g.add((this_sample, PROV.wasAttributedTo, ga))
	// g.add((ga, RDFS.label, Literal('Geoscience Australia', datatype=XSD.string)))

	// Add in
	newsub, _ := rdf.NewIRI("http://foo.org/thisSample")
	newpred1, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	newobj1, _ := rdf.NewIRI("http://www.w3.org/ns/prov#entity")
	tr = append(tr, rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1})

	ga, _ := rdf.NewIRI("http://opencoredata.org/org") // ?
	newpred2, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	newobj2, _ := rdf.NewIRI("http://www.w3.org/ns/prov#org")
	tr = append(tr, rdf.Triple{Subj: ga, Pred: newpred2, Obj: newobj2})

	bn, _ := rdf.NewBlank("bn1")

	newpred3, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	newobj3, _ := rdf.NewIRI("http://www.w3.org/ns/prov#attribution")
	tr = append(tr, rdf.Triple{Subj: bn, Pred: newpred3, Obj: newobj3})

	newpred4, _ := rdf.NewIRI("http://www.w3.org/ns/prov#agent")
	tr = append(tr, rdf.Triple{Subj: bn, Pred: newpred4, Obj: ga})

	newpred5, _ := rdf.NewIRI("http://www.w3.org/ns/prov#hadRole")
	newobj5, _ := rdf.NewIRI("http://www.aurole.org/Publisher")
	tr = append(tr, rdf.Triple{Subj: bn, Pred: newpred5, Obj: newobj5})

	newpred6, _ := rdf.NewIRI("http://www.w3.org/ns/prov#qualifiedAttribution")
	tr = append(tr, rdf.Triple{Subj: newsub, Pred: newpred6, Obj: bn})

	newpred7, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#type")
	newobj7, _ := rdf.NewIRI("http://www.w3.org/ns/prov#Agent")
	tr = append(tr, rdf.Triple{Subj: ga, Pred: newpred7, Obj: newobj7})

	newpred8, _ := rdf.NewIRI("http://www.w3.org/ns/prov#wasAttributedTo")
	tr = append(tr, rdf.Triple{Subj: newsub, Pred: newpred8, Obj: bn})

	newpred9, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#label")
	newobj9, _ := rdf.NewLiteral("Geoscience Australia")
	tr = append(tr, rdf.Triple{Subj: ga, Pred: newpred9, Obj: newobj9})

	fmt.Println(tr)

	var inoutFormat rdf.Format
	inoutFormat = rdf.Turtle //NTriples

	// Create a buffer io writer
	// var b bytes.Buffer
	// foo := bufio.NewWriter(&b)

	// Create output file
	outFile, err := os.Create("prov.nt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	enc := rdf.NewTripleEncoder(outFile, inoutFormat)
	err = enc.EncodeAll(tr)
	err = enc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// foo.Flush()
	// fmt.Println(string(b.Bytes()))
}
