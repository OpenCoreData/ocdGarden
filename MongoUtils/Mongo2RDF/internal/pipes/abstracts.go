package pipes

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/knakk/rdf"
	mgo "gopkg.in/mgo.v2"
	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/common"
	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/structs"
)

// NewAbstracts update
func NewAbstracts(session *mgo.Session, ub *common.Buffer) int {

	var b strings.Builder

	// Make a context for this graph
	newctx, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/objectgraph/id/%s", "CSDCOABSID"))
	ctx := rdf.Context(newctx)

	// connect and get documents
	csvw := session.DB("abstracts").C("csdco")
	var csdcoAbs []ocdstructs.MdocsV2
	err := csvw.Find(nil).All(&csdcoAbs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	for _, item := range csdcoAbs {
		abstractIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s", item.ID)
		if item.Type != "" {
			// _ = common.ILTriple(abstractIRI, "http://opencoredata.org/voc/janus/v1/Type", item.Type, ctx, &b)
			_ = common.IITriple(abstractIRI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://opencoredata.org/voc/csdco/v1/Abstract", ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/title", common.StripCtlAndExtFromUnicode(item.Title), ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/type", item.Type, ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/year", strconv.Itoa(item.Year), ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/source", item.Source, ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/id", item.ID, ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/created", item.Created.String(), ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/profileid", item.Profile_ID, ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/groupid", item.Group_ID, ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/lastmodified", item.Last_Modified.String(), ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/abstracttext", common.StripCtlAndExtFromUnicode(item.Abstract), ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/doi", item.Identifiers.Doi, ctx, &b)
			_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/issn", item.Identifiers.Issn, ctx, &b)
			// loop on Tags
			if len(item.Tags) > 0 {
				tagIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s#tag", item.ID)
				_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/tag", tagIRI, ctx, &b)

				for _, tag := range item.Tags {
					_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/tag/value", common.StripCtlAndExtFromUnicode(tag), ctx, &b)
				}
			}

			// loop on Authors
			if len(item.Authors) > 0 {
				authorIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s#authors", item.ID)
				_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/authors", authorIRI, ctx, &b)

				for _, author := range item.Authors {
					_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/author/firstname", common.StripCtlAndExtFromUnicode(author.First_Name), ctx, &b)
					_ = common.ILTriple(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/author/lastname", common.StripCtlAndExtFromUnicode(author.Last_Name), ctx, &b)
				}
			}
		}
	}

	len, err := ub.Write([]byte(b.String()))
	if err != nil {
		log.Printf("error in the buffer write... %v\n", err)
	}

	return len
}

// Abstracts CSDCO abstract functions
func Abstracts(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("abstracts").C("csdco")
	var csdcoAbs []ocdstructs.MdocsV2
	err := csvw.Find(nil).All(&csdcoAbs)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range csdcoAbs {
		// Make subject IRI
		// newsub, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/id/resource/janus/query/%s", item.ID)) // Sprintf a correct URI here

		// title
		// newpred0, _ := rdf.NewIRI("http://www.w3.org/1999/02/22-rdf-syntax-ns#title")
		// newobj0, _ := rdf.NewIRI("http://opencoredata.org/id/voc/janus/v1/JanusQuery")
		// newtriple0 := rdf.Triple{Subj: newsub, Pred: newpred0, Obj: newobj0}
		// tr = append(tr, newtriple0)

		abstractIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s", item.ID)

		// need a TYPE!
		common.SPOIRI(abstractIRI, "a", "http://opencoredata.org/id/voc/csdco/v1/Abstract")

		if item.Title != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/title", common.StripCtlAndExtFromUnicode(item.Title)))
		}
		if item.Type != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/type", item.Type))
		}
		if item.Year != 0 {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/year", strconv.Itoa(item.Year)))
		}
		if item.Source != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/source", item.Source))
		}
		if item.ID != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/id", item.ID))
		}
		if !item.Created.IsZero() {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/created", item.Created.String()))
		}
		if item.Profile_ID != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/profileid", item.Profile_ID))
		}
		if item.Group_ID != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/groupid", item.Group_ID))
		}
		if !item.Last_Modified.IsZero() {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/lastmodified", item.Last_Modified.String()))
		}
		if item.Abstract != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/abstracttext", common.StripCtlAndExtFromUnicode(item.Abstract)))
		}
		if item.Identifiers.Doi != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/doi", item.Identifiers.Doi))
		}
		if item.Identifiers.Issn != "" {
			tr = append(tr, common.SPOLiteral(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/issn", item.Identifiers.Issn))
		}

		// loop on Tags
		if len(item.Tags) > 0 {
			tagIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s#tag", item.ID)
			tr = append(tr, common.SPOIRI(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/tag", tagIRI))
			for _, tag := range item.Tags {
				tr = append(tr, common.SPOLiteral(tagIRI, "http://opencoredata.org/id/voc/csdco/v1/tag/value", common.StripCtlAndExtFromUnicode(tag)))
			}
		}

		// loop on Authors
		if len(item.Authors) > 0 {
			authorIRI := fmt.Sprintf("http://opencoredata.org/id/resource/csdco/abstract/%s#authors", item.ID)
			tr = append(tr, common.SPOIRI(abstractIRI, "http://opencoredata.org/id/voc/csdco/v1/authors", authorIRI))
			for _, author := range item.Authors {
				tr = append(tr, common.SPOLiteral(authorIRI, "http://opencoredata.org/id/voc/csdco/v1/author/firstname", common.StripCtlAndExtFromUnicode(author.First_Name)))
				tr = append(tr, common.SPOLiteral(authorIRI, "http://opencoredata.org/id/voc/csdco/v1/author/lastname", common.StripCtlAndExtFromUnicode(author.Last_Name)))
			}
		}

		// for _, author := range item.Authors {
		// 	fmt.Printf("FirstName: %s\n", author.First_Name)
		// 	fmt.Printf("LastName: %s\n", author.Last_Name)
		// }

	}

	common.WriteFile("./output/csdcoAbstracts.nt", tr)

}
