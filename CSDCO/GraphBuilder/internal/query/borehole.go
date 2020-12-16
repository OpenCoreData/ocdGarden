package query

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	"github.com/minio/minio-go"
	gominio "github.com/minio/minio-go"

	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/internal/jld"
)

// Boreholes will query for the projects and print the resulted
func Boreholes(db *sql.DB, mc *minio.Client, bucketVal string) {
	rows, err := db.Query(`SELECT  Azimuth,  Dip,   Elevation, Lat, Long,   Water_Depth,   mblf_B, mblf_T, Site,
	Comment, Country, County_Region, Date, Expedition, Hole, Hole_ID, IGSN, Location, Location_ID, Location_Type,
	Metadata_Source, NGDC_Serial, Original_ID, PI, Platform, Position, Sample_Type, SiteHole, State_Province,
	platform_name, platform_type FROM boreholes`)
	if err != nil {
		log.Println(err)
	}

	//bucketVal := "csdco"

	for rows.Next() {
		var Azimuth, Dip, Elevation, Lat, Long, Water_Depth, mblf_B, mblf_T sql.NullFloat64
		var Site sql.NullInt64
		var Comment, Country, County_Region, Date, Expedition, Hole, Hole_ID, IGSN, Location, Location_ID, Location_Type,
			Metadata_Source, NGDC_Serial, Original_ID, PI, Platform, Position, Sample_Type, SiteHole,
			State_Province, platform_name, platform_type sql.NullString

		err := rows.Scan(&Azimuth, &Dip, &Elevation, &Lat, &Long, &Water_Depth, &mblf_B, &mblf_T, &Site,
			&Comment, &Country, &County_Region, &Date, &Expedition, &Hole, &Hole_ID, &IGSN, &Location, &Location_ID, &Location_Type,
			&Metadata_Source, &NGDC_Serial, &Original_ID, &PI, &Platform, &Position, &Sample_Type, &SiteHole,
			&State_Province, &platform_name, &platform_type)
		if err != nil {
			log.Println(err)
		}

		p := jld.CSDCOBorehole{Azimuth.Float64, Dip.Float64, Elevation.Float64, Lat.Float64, Long.Float64, Water_Depth.Float64,
			mblf_B.Float64, mblf_T.Float64, Site.Int64,
			Comment.String, Country.String, County_Region.String, Date.String, Expedition.String, Hole.String, Hole_ID.String, IGSN.String, Location.String, Location_ID.String, Location_Type.String,
			Metadata_Source.String, NGDC_Serial.String, Original_ID.String, PI.String, Platform.String, Position.String, Sample_Type.String, SiteHole.String,
			State_Province.String, platform_name.String, platform_type.String}

		jld, err := jld.BoreholeDG(p)
		if err != nil {
			log.Println(err)
		}

		// fmt.Println(string(jld))
		fmt.Printf("Boreholes: %s in %s\n", Hole_ID.String, fmt.Sprintf("%s-do-resources", bucketVal))

		b := bytes.NewBuffer(jld)

		contentType := "application/ld+json" // really Nq right now
		//n, err := mc.PutObject("doa-meta", objectName, b, int64(b.Len()), minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})
		_, err = mc.PutObject(fmt.Sprintf("%s-do-resources", bucketVal), fmt.Sprintf("%s", Hole_ID.String),
			b, int64(b.Len()), gominio.PutObjectOptions{ContentType: contentType})
		// log.Printf("Loading metadata object: %d\n", n)  // was printing the returned byte length from PutObject
		if err != nil {
			log.Printf("Error loading metadata object to minio bucket %s : %s\n", bucketVal, err)
		}
	}
	rows.Close() //good habit to close
}
