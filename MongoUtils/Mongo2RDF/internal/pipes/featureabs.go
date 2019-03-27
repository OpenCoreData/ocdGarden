package pipes

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/knakk/rdf"
	mgo "gopkg.in/mgo.v2"
	"opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/common"
	ocdstructs "opencoredata.org/ocdGarden/MongoUtils/Mongo2RDF/internal/structs"
)

// NewFeatureAbsGeoJSON makes JRSO festure graph
func NewFeatureAbsGeoJSON(session *mgo.Session, ub *common.Buffer) int {
	var b strings.Builder

	// Make a context for this graph
	newctx, _ := rdf.NewIRI(fmt.Sprintf("http://opencoredata.org/objectgraph/id/%s", "jrsoexp"))
	ctx := rdf.Context(newctx)

	// connect and get documents
	csvw := session.DB("expedire").C("featuresAbsGeoJSON")
	var features []ocdstructs.ExpeditionGeoJSON
	err := csvw.Find(nil).All(&features)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	for _, item := range features {
		if item.Type != "" {

			// TODO  when does item.Type be anything but 1 thing?
			// split item.Uri into LegURI   SiteURI and HoleURI
			s := item.Uri
			u, err := url.Parse(s)
			if err != nil {
				log.Println(item.Uri)
				log.Println(err)

			}
			pa := strings.Split(u.Path, "/")

			// log.Println(pa)
			legURI := fmt.Sprintf("http://opencoredata.org/id/expedition/%s", pa[3])
			siteURI := fmt.Sprintf("http://opencoredata.org/id/expedition/%s/%s", pa[3], pa[4])
			holeURI := fmt.Sprintf("http://opencoredata.org/id/expedition/%s/%s/%s", pa[3], pa[4], "A") // not all legs have holes..  only leg site..   so give them a "hole" A
			if len(pa) > 5 {
				holeURI = fmt.Sprintf("http://opencoredata.org/id/expedition/%s/%s/%s", pa[3], pa[4], pa[5])
			}

			// type Leg
			_ = common.IITriple(legURI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://opencoredata.org/voc/janus/v1/Leg", ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/expeditionlabel", item.Expedition, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/note", item.Note, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/prcoeedingreport", item.Prcoeedingreport, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/abstract", common.StripCtlAndExtFromUnicode(item.Abstract), ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/program", item.Program, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/initialreportvolume", item.Initialreportvolume, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/scientificprospectus", item.Scientificprospectus, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/scientificreportvolume", item.Scientificreportvolume, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/expeditionsitelabel", item.Expeditionsite, ctx, &b)
			_ = common.ILTriple(legURI, "http://opencoredata.org/voc/janus/v1/preliminaryreport", item.Preliminaryreport, ctx, &b)
			_ = common.IITriple(legURI, "http://opencoredata.org/voc/janus/v1/hasSite", siteURI, ctx, &b)
			_ = common.IITriple(legURI, "http://opencoredata.org/voc/janus/v1/hasHole", holeURI, ctx, &b)

			// type Site
			_ = common.IITriple(siteURI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://opencoredata.org/voc/janus/v1/Site", ctx, &b)
			_ = common.ILTriple(siteURI, "http://opencoredata.org/voc/janus/v1/sitelabel", item.Site, ctx, &b)
			_ = common.IITriple(siteURI, "http://opencoredata.org/voc/janus/v1/hasHole", holeURI, ctx, &b)

			// type Hole
			_ = common.IITriple(holeURI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type", "http://opencoredata.org/voc/janus/v1/Hole", ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/holelabel", item.Hole, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/coreinterval", item.CoreInterval, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/percentrecovery", item.PercentRecovery, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/drilled", item.Drilled, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/vcdata", item.Vcdata, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/waterdepth", item.Waterdepth, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/corecount", item.CoreCount, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/coredata", item.Coredata, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/logdata", item.Logdata, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/geom", item.Geom, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/corerecovery", item.CoreRecovery, ctx, &b)
			_ = common.ILTriple(holeURI, "http://opencoredata.org/voc/janus/v1/penetration", item.Penetration, ctx, &b)
		}
	}

	len, err := ub.Write([]byte(b.String()))
	if err != nil {
		log.Printf("error in the buffer write... %v\n", err)
	}

	return len
}

//DELETEFeatureAbsGeoJSON  to be removed
func DELETEFeatureAbsGeoJSON(session *mgo.Session) {

	// connect and get documents
	csvw := session.DB("expedire").C("featuresAbsGeoJSON")
	var features []ocdstructs.ExpeditionGeoJSON
	err := csvw.Find(nil).All(&features)
	if err != nil {
		fmt.Printf("this is error %v \n", err)
	}

	// RDF item
	tr := []rdf.Triple{}

	// Loop on documents
	for _, item := range features {

		common.SPOIRI(item.Uri, "a", "http://opencoredata.org/voc/janus/v1/Feature")

		if item.Type != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Type", item.Type))
		}
		if item.Hole != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Hole", item.Hole))
		}
		if item.Expedition != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Expedition", item.Expedition))
		}
		if item.Site != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Site", item.Site))
		}
		if item.Program != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Program", item.Program))
		}
		if item.Waterdepth != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Waterdepth", item.Waterdepth))
		}
		if item.CoreCount != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Corecount", item.CoreCount))
		}
		if item.Initialreportvolume != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Initialreportvolume", item.Initialreportvolume))
		}
		if item.Coredata != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Coredata", item.Coredata))
		}
		if item.Logdata != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Logdata", item.Logdata))
		}
		if item.Geom != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Geom", item.Geom))
		}
		if item.Scientificprospectus != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Scientificprospectus", item.Scientificprospectus))
		}
		if item.CoreRecovery != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Corerecovery", item.CoreRecovery))
		}
		if item.Penetration != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Penetration", item.Penetration))
		}
		if item.Scientificreportvolume != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Scientificreportvolume", item.Scientificreportvolume))
		}
		if item.Expeditionsite != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Expeditionsite", item.Expeditionsite))
		}
		if item.Preliminaryreport != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Preliminaryreport", item.Preliminaryreport))
		}
		if item.CoreInterval != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Coreinterval", item.CoreInterval))
		}
		if item.PercentRecovery != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Percentrecovery", item.PercentRecovery))
		}
		if item.Drilled != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Drilled", item.Drilled))
		}
		if item.Vcdata != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/VCdata", item.Vcdata))
		}
		if item.Note != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Note", item.Note))
		}
		if item.Prcoeedingreport != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Prcoeedingreport", item.Prcoeedingreport))
		}
		if item.Abstract != "" {
			tr = append(tr, common.SPOLiteral(item.Uri, "http://opencoredata.org/voc/janus/v1/Abstract", common.StripCtlAndExtFromUnicode(item.Abstract)))
		}

	}

	common.WriteFile("./output/janusAbstracts.nt", tr)

}
