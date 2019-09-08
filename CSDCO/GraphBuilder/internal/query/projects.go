package query

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	gominio "github.com/minio/minio-go"

	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/internal/connectors"
	"opencoredata.org/ocdGarden/CSDCO/GraphBuilder/internal/jld"
)

// Projects query
func Projects(db *sql.DB) {
	rows, err := db.Query(`SELECT Expedition,Full_Name,Funding,Technique,Discipline,
	Link_Title,Link_URL,Lab,Repository,Status,Start_Date,Outreach,Investigators,Abstract  FROM projects`)
	if err != nil {
		log.Println(err)
	}

	mc := connectors.MinioConnection()
	bucketVal := "csdco"

	for rows.Next() {
		var Expedition, Full_Name, Funding, Technique, Discipline, Link_Title, Link_URL,
			Lab, Repository, Status, Start_Date, Outreach, Investigators, Abstract sql.NullString

		err = rows.Scan(&Expedition, &Full_Name, &Funding, &Technique, &Discipline,
			&Link_Title, &Link_URL, &Lab, &Repository, &Status, &Start_Date, &Outreach, &Investigators, &Abstract)
		if err != nil {
			log.Println(err)
		}

		pf := []jld.PjctFeature{}

		var Hole_ID, IGSN, Location, Location_ID sql.NullString
		var Lat, Long sql.NullFloat64

		r, err := db.Query(`SELECT   Lat, Long, Hole_ID, IGSN, Location,
		 Location_ID FROM boreholes WHERE Expedition is ?`, Expedition)
		if err != nil {
			log.Println(err)
		}

		for r.Next() {
			err = r.Scan(&Lat, &Long, &Hole_ID, &IGSN, &Location, &Location_ID)
			if err != nil {
				log.Println(err)
			}

			f := jld.PjctFeature{Name: Hole_ID.String, Latitude: Lat.Float64, Longitude: Long.Float64}
			pf = append(pf, f)
		}
		r.Close()

		p := jld.Project{Expedition: Expedition.String, FullName: Full_Name.String, Funding: Funding.String,
			Technique: Technique.String, Discipline: Discipline.String, LinkTitle: Link_Title.String,
			LinkURL: Link_URL.String, Lab: Lab.String, Repository: Repository.String, Status: Status.String,
			StartDate: Start_Date.String, Outreach: Outreach.String,
			Investigators: Investigators.String, Abstract: Abstract.String, Features: pf}

		jld, err := jld.ProjectDG(p)
		if err != nil {
			log.Println(err)
		}

		// fmt.Println(string(jld))
		fmt.Printf("Project: %s in %s\n", Expedition.String, fmt.Sprintf("%s-do-resources", bucketVal))

		// load to minio with an ID for the object (sha256)
		// b := bytes.NewBufferString(lb.String())  // when sending NQ, convert the string to a io reader bytes buffer string
		b := bytes.NewBuffer(jld) // if conversting lb to JSON-LD then that comes back as byte array, so make a new byte buffer

		contentType := "application/ld+json" // really Nq right now
		//n, err := mc.PutObject("doa-meta", objectName, b, int64(b.Len()), minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})
		_, err = mc.PutObject(fmt.Sprintf("%s-do-resources", bucketVal), fmt.Sprintf("%s", Expedition.String),
			b, int64(b.Len()), gominio.PutObjectOptions{ContentType: contentType})
		// log.Printf("Loading metadata object: %d\n", n)  // was printing the returned byte length from PutObject
		if err != nil {
			log.Printf("Error loading metadata object to minio bucket %s : %s\n", bucketVal, err)
		}
	}
	rows.Close() //good habit to close
}
