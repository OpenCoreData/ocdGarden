package report

import (
	"fmt"
	"log"
	"strings"

	"../vault"
	"github.com/knakk/rdf"
	"github.com/rs/xid"

	"../../pkg/utils"
)

// RDFGraph (item, shaval, *rdf)
// In this approach each object gets a named graph.  Perhaps this is not
// needed since each data graph also has a sha ID with it?  Whic is all we really
// use in the graph IRI.   ???
func RDFGraph(item vault.VaultItem, shaval string, ub *utils.Buffer) int {

	t := utils.MimeByType(item.FileExt)

	newctx, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/objectgraph/id/%s", shaval))
	ctx := rdf.Context(newctx)

	guid := xid.New()
	newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/id/do/%s", guid)) // Sprintf a correct URI here

	// need rdf:type triples  (include dcat dataset)
	newpred7, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#") // dcat:type -> multitype schema.org/Dataset  opencoredata.org/X
	newobj7, _ := rdf.NewIRI(item.TypeURI)                                   //  http://www.w3.org/ns/dcat#   Dataset   title  keywords
	newtriple7 := rdf.Triple{Subj: newsub, Pred: newpred7, Obj: newobj7}
	nq7 := rdf.Quad{newtriple7, ctx}

	// review the schema.org type dataset for properties

	newpred1, _ := rdf.NewIRI("http://opencoredata.org/voc/csdco/v1/Project")
	newobj1, _ := rdf.NewLiteral(item.Project)
	newtriple1 := rdf.Triple{Subj: newsub, Pred: newpred1, Obj: newobj1}
	nq1 := rdf.Quad{newtriple1, ctx}

	newpred2, _ := rdf.NewIRI("http://opencoredata.org/voc/csdco/v1/Type") // dcat:description make this a literal type definition
	newobj2, _ := rdf.NewLiteral(item.Type)
	newtriple2 := rdf.Triple{Subj: newsub, Pred: newpred2, Obj: newobj2}
	nq2 := rdf.Quad{newtriple2, ctx}

	newpred3, _ := rdf.NewIRI("http://opencoredata.org/voc/csdco/v1/FileName") //Distribution:title
	newobj3, _ := rdf.NewLiteral(item.FileName)
	newtriple3 := rdf.Triple{Subj: newsub, Pred: newpred3, Obj: newobj3}
	nq3 := rdf.Quad{newtriple3, ctx}

	newpred4, _ := rdf.NewIRI("http://opencoredata.org/voc/csdco/v1/FileExt")
	newobj4, _ := rdf.NewLiteral(item.FileExt)
	newtriple4 := rdf.Triple{Subj: newsub, Pred: newpred4, Obj: newobj4}
	nq4 := rdf.Quad{newtriple4, ctx}

	newpred5, _ := rdf.NewIRI("http://opencoredata.org/voc/csdco/v1/SHAHash")
	newobj5, _ := rdf.NewLiteral(shaval)
	newtriple5 := rdf.Triple{Subj: newsub, Pred: newpred5, Obj: newobj5}
	nq5 := rdf.Quad{newtriple5, ctx}

	newpred6, _ := rdf.NewIRI("http://opencoredata.org/voc/csdco/v1/Mime") // Distribution: mediatype
	newobj6, _ := rdf.NewLiteral(t)
	newtriple6 := rdf.Triple{Subj: newsub, Pred: newpred6, Obj: newobj6}
	nq6 := rdf.Quad{newtriple6, ctx}

	// serialize....
	// since this can be serilized to string, rather than an array of structs
	// just make a string buffer and append the serialized triples to it..
	// then return these..  allows easy use of a thread safe bytes buffer then.
	//qd := []rdf.Quad{}
	var b strings.Builder

	if newtriple1.Obj.String() != "" {
		//	qd = append(qd, nq1)
		s := nq1.Serialize(rdf.NQuads)
		fmt.Fprintf(&b, "%s", s)
	}

	if newtriple2.Obj.String() != "" {
		//	qd = append(qd, nq2)
		s := nq2.Serialize(rdf.NQuads)
		fmt.Fprintf(&b, "%s", s)
	}

	if newtriple3.Obj.String() != "" {
		//	qd = append(qd, nq3)
		s := nq3.Serialize(rdf.NQuads)
		fmt.Fprintf(&b, "%s", s)
	}

	if newtriple4.Obj.String() != "" {
		//	qd = append(qd, nq4)
		s := nq4.Serialize(rdf.NQuads)
		fmt.Fprintf(&b, "%s", s)
	}

	if newtriple5.Obj.String() != "" {
		//	qd = append(qd, nq5)
		s := nq5.Serialize(rdf.NQuads)
		fmt.Fprintf(&b, "%s", s)
	}

	if newtriple6.Obj.String() != "" {
		//	qd = append(qd, nq6)
		s := nq6.Serialize(rdf.NQuads)
		fmt.Fprintf(&b, "%s", s)
	}

	if newtriple7.Obj.String() != "" {
		//	qd = append(qd, nq7)
		s := nq7.Serialize(rdf.NQuads)
		fmt.Fprintf(&b, "%s", s)
	}

	len, err := ub.Write([]byte(b.String()))
	if err != nil {
		log.Printf("error in the buffer write... %v\n", err)
	}

	return len //  we will return the bytes count we write...

	//return []byte(b.String())
	// TODO return this as a []byte and we can append this to our thread safe buffer
	//  look at millers graph for an implementation for main
}
