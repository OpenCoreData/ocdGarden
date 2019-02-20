package fetch

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"earthcube.org/Project418/gleaner/pkg/summoner/sitemaps"
	"earthcube.org/Project418/gleaner/pkg/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
	minio "github.com/minio/minio-go"

	localutils "opencoredata.org/ocdGarden/janusPackageBuilder/internal/utils"
)

func ReadSitemap(mc *minio.Client) error {
	cs := utils.Config{}
	us := sitemaps.IngestSitemapXML("http://opencoredata.org/sitemap.xml", cs)

	for i := range us.URL {
		// fmt.Println(us.URL[i].Loc)
		// fmt.Println(urlUUID(us.URL[i].Loc))

		j, err := jld(us.URL[i].Loc)
		if err != nil {
			log.Println(err)
		}

		d, err := dcURL(j)
		if err != nil {
			log.Println(err)
		}

		b, err := dcBody(d)
		if err != nil {
			log.Println(err)
		}

		// now load j and b into Minio
		// err := osload("bucket", "objectname", "content")
		if err == nil {
			fmt.Printf("Writing %s to minio\n", urlUUID(us.URL[i].Loc))
			localutils.LoadToMinio(j, "jsdo", urlUUID(us.URL[i].Loc), mc)
			localutils.LoadToMinio(string(b), "jdu", urlUUID(us.URL[i].Loc), mc)
		}

	}

	return nil
}

// use gleaner code and place in pkg there...
// might become a package there.   just some code here...
func jld(url string) (string, error) {

	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// not even being able to make a req instance..  might be a fatal thing?
		log.Printf("------ error making request------ \n %s", err)
	}

	req.Header.Set("User-Agent", "EarthCube_DataBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error reading location: %s", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Printf("Error doc from resp: %v", err)
	}

	var jsonld string
	if err == nil {
		doc.Find("script").Each(func(i int, s *goquery.Selection) {
			val, _ := s.Attr("type")
			if val == "application/ld+json" {
				// err = isValid(s.Text())
				// if err != nil {
				// 		log.Printf("ERROR: At %s JSON-LD is NOT valid: %s", urlloc, err)
				//	}
				jsonld = s.Text()
			}
		})
	}

	return jsonld, err
}

// use gleaner code and place in pkg there...
// I must do something like this in the Tika code...
func dcURL(jsonld string) (string, error) {
	dl, err := jsonparser.GetString([]byte(jsonld), "distribution", "contentUrl")
	if err != nil {
		log.Println(err)
		return "", err
	}

	// TODO   convert back to []byte and reutrn the dl content..  not the URL
	return dl, err
}

func dcBody(url string) ([]byte, error) {
	rd, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rd.Body.Close()
	r, _ := ioutil.ReadAll(rd.Body)

	return r, err
}

func urlUUID(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	p := strings.Split(u.Path, "/")
	return p[len(p)-1]
}
