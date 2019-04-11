package report

import (
	"fmt"
	"log"
	"strings"

	"github.com/knakk/rdf"
	"github.com/rs/xid"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/vault"

	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"
)

// TODO
// schema.org/DataDownload

// RDFGraph (item, shaval, *rdf)
// In this approach each object gets a named graph.  Perhaps this is not
// needed since each data graph also has a sha ID with it?  Which is all we really
// use in the graph IRI.   ???
func RDFGraph(item vault.VaultItem, shaval string, ub *utils.Buffer) int {
	var b strings.Builder

	t := utils.MimeByType(item.FileExt)
	newctx, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/objectgraph/id/%s", shaval))
	ctx := rdf.Context(newctx)

	guid := xid.New()
	s := fmt.Sprintf("http://opencoredata.org/id/do/%s", guid)
	d := fmt.Sprintf("http://opencoredata.org/id/dx/%s", guid) // distribution URL

	_ = iiTriple(s, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", item.TypeURI, ctx, &b)
	_ = iiTriple(s, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://www.w3.org/ns/dcat#Dataset", ctx, &b)
	_ = iiTriple(s, "http://purl.org/dc/terms/type", "http://purl.org/dc/dcmitype/Dataset", ctx, &b)
	_ = iiTriple(s, "http://www.w3.org/ns/dcat#distribution", d, ctx, &b)
	//  _ = iiTriple(s, "", "", ctx, &b)  //  should we add in a landing page?
	// _ = iiTriple(s, "http://www.w3.org/2000/01/rdf-schema#seeAlso". "cruise URI" )
	//  If there is an inbox, would need to look here...   (some generic manner to do this?)

	_ = ilTriple(s, "http://opencoredata.org/voc/csdco/v1/Project", item.Project, ctx, &b)
	_ = ilTriple(s, "http://opencoredata.org/voc/csdco/v1/Type", item.Type, ctx, &b)
	_ = ilTriple(s, "http://opencoredata.org/voc/csdco/v1/FileName", item.FileName, ctx, &b)
	_ = ilTriple(s, "http://opencoredata.org/voc/csdco/v1/FileAge", fmt.Sprintf("%f", item.Age), ctx, &b)
	_ = ilTriple(s, "http://opencoredata.org/voc/csdco/v1/FileExt", item.FileExt, ctx, &b)
	_ = ilTriple(s, "http://opencoredata.org/voc/csdco/v1/SHAHash", shaval, ctx, &b)
	_ = ilTriple(s, "http://opencoredata.org/voc/csdco/v1/Mime", t, ctx, &b)

	_ = iiTriple(d, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://www.w3.org/ns/dcat#Distribution", ctx, &b)
	_ = iiTriple(d, "http://purl.org/dc/terms/license", "http://example.com/cc0.html", ctx, &b)
	_ = iiTriple(d, "http://www.w3.org/ns/dcat#downloadURL", s, ctx, &b)
	// _ = iiTriple(d, "http://www.w3.org/ns/dcat#mediaType", "https://www.iana.org/assignments/media-types/text/csv", ctx, &b)

	_ = ilTriple(d, "http://purl.org/dc/terms/title", fmt.Sprintf("Digital object %s for CSDCO project %s", item.FileName, item.Project), ctx, &b)
	// _ = ilTriple(d, "http://purl.org/dc/terms/description", "Description info here", ctx, &b)

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
