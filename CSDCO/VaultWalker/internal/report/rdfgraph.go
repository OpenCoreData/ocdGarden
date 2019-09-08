package report

import (
	"fmt"
	"log"
	"strings"

	"github.com/knakk/rdf"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/vault"

	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"
)

// TODO
// schema.org/DataDownload

// RDFGraph (item, shaval, *rdf)
// In this approach each object gets a named graph.  Perhaps this is not
// needed since each data graph also has a sha ID with it?  Which is all we really
// use in the graph IRI.   ???
func RDFGraph(guid string, item vault.VaultItem, shaval string, ub *utils.Buffer) int {
	var b strings.Builder

	t := utils.MimeByType(item.FileExt)
	newctx, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/objectgraph/id/%s", shaval))
	ctx := rdf.Context(newctx)

	douri := fmt.Sprintf("http://opencoredata.org/id/do/%s", shaval) // before I was using the guid var value..  not sure why

	// guid := xid.New()                                          // Not sure why I was using xid here..
	s := fmt.Sprintf("http://opencoredata.org/id/do/%s", guid)
	// before I was using the guid var value..  not sure why
	// answer?  I may add or remove triples..  so the hash is a bad ID here.
	// it's fine when I'm dealing with a byte stream and that is likely immutable

	bn, err := rdf.NewBlank(fmt.Sprintf("b%s", guid))
	if err != nil {
		log.Println("Error setting blank node")
	}

	_ = iiTriple(s, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://www.schema.org/DigitalDocument", ctx, &b)
	_ = ilTriple(s, "http://schema.org/description", fmt.Sprintf("Digital object of type %s named %s for CSDCO project %s", item.Type, item.FileName, item.Project), ctx, &b)
	_ = ilTriple(s, "http://schema.org/isRelatedTo", item.Project, ctx, &b)
	_ = ilTriple(s, "http://schema.org/name", item.FileName, ctx, &b)
	_ = ilTriple(s, "http://schema.org/dateCreated", item.DateCreated, ctx, &b)

	_ = ilTriple(s, "http://schema.org/encodingFormat", t, ctx, &b)
	_ = iiTriple(s, "http://schema.org/additionalType", item.TypeURI, ctx, &b) // should be the URL
	_ = iiTriple(s, "http://schema.org/license", "https://creativecommons.org/share-your-work/public-domain/cc0/", ctx, &b)
	_ = iiTriple(s, "http://schema.org/url", douri, ctx, &b)
	_ = ibTriple(s, "http://schema.org/identifier", bn, ctx, &b)

	_ = blTriple(bn, "http://schema.org/propertyID", "SHA256", ctx, &b)
	_ = blTriple(bn, "http://schema.org/value", shaval, ctx, &b)
	_ = biTriple(bn, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://schema.org/PropertyValue", ctx, &b)

	len, err := ub.Write([]byte(b.String()))
	if err != nil {
		log.Printf("error in the buffer write... %v\n", err)
	}

	return len //  we will return the bytes count we write...
}

func iiTriple(s, p, o string, c rdf.Context, b *strings.Builder) error {
	sub, err := rdf.NewIRI(s)
	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewIRI(o)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	if s != "" && p != "" && o != "" {
		fmt.Fprintf(b, "%s", qs)
	}

	return err
}

func ilTriple(s, p, o string, c rdf.Context, b *strings.Builder) error {
	sub, err := rdf.NewIRI(s)
	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewLiteral(o)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	if s != "" && p != "" && o != "" {
		fmt.Fprintf(b, "%s", qs)
	}

	return err
}

func ibTriple(s, p string, o rdf.Blank, c rdf.Context, b *strings.Builder) error {
	sub, err := rdf.NewIRI(s)
	pred, err := rdf.NewIRI(p)

	t := rdf.Triple{Subj: sub, Pred: pred, Obj: o}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	if s != "" && p != "" {
		fmt.Fprintf(b, "%s", qs)
	}

	return err
}

func blTriple(s rdf.Blank, p, o string, c rdf.Context, b *strings.Builder) error {

	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewLiteral(o)

	t := rdf.Triple{Subj: s, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	if p != "" && o != "" {
		fmt.Fprintf(b, "%s", qs)
	}

	return err
}

func biTriple(s rdf.Blank, p, o string, c rdf.Context, b *strings.Builder) error {

	pred, err := rdf.NewIRI(p)
	obj, err := rdf.NewIRI(o)

	t := rdf.Triple{Subj: s, Pred: pred, Obj: obj}
	q := rdf.Quad{t, c}

	qs := q.Serialize(rdf.NQuads)
	if p != "" && o != "" {
		fmt.Fprintf(b, "%s", qs)
	}

	return err
}
