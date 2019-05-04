package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	gominio "github.com/minio/minio-go"
	"github.com/rs/xid"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/index"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/minio"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/report"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/vault"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"
)

var minioVal, portVal, accessVal, secretVal, bucketVal, dirVal string
var uploadVal, sslVal bool

func init() {
	akey := os.Getenv("MINIO_ACCESS_KEY")
	skey := os.Getenv("MINIO_SECRET_KEY")
	addr := os.Getenv("MINIO_ADDRESS")
	port := os.Getenv("MINIO_PORT")
	bckt := os.Getenv("MINIO_BUCKET")
	// ssl := os.Getenv("MINIO_SSL")

	flag.StringVar(&minioVal, "address", addr, "FQDN for server")
	flag.StringVar(&portVal, "port", port, "Port for minio server, default 9000")
	flag.StringVar(&accessVal, "access", akey, "Access Key ID")
	flag.StringVar(&secretVal, "secret", skey, "Secret access key")
	flag.StringVar(&bucketVal, "bucket", bckt, "The configuration bucket")
	flag.BoolVar(&sslVal, "ssl", false, "Use SSL boolean")

	flag.BoolVar(&uploadVal, "upload", false, "Upload files to object store")
	flag.StringVar(&dirVal, "d", "./test", "Directory to walk")
}

func main() {
	// Set up some vars..  parse the flags and get a minio connection
	start := time.Now()
	var files []string
	var va []vault.VaultItem
	flag.Parse()
	mc := minio.MinioConnection(minioVal, portVal, accessVal, secretVal)

	// Get directory and walk it and put a paths in files string array
	d := dirVal
	err := filepath.Walk(d, visit(&files))
	if err != nil {
		panic(err)
	}

	// make output directory if it doesn't exist
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.Mkdir("./output", os.ModePerm)
	}

	// get the various elements of the found files, walk the files string array
	for _, file := range files {
		v, err := index.PathInspection(d, file)
		if err == nil {
			va = append(va, v)
		}
	}

	vh := vault.VaultHoldings{va}
	pl := vh.Prjs() // get unique project (dir) names

	fmt.Printf("{Projects: %s\n", pl)
	var b utils.Buffer

	for _, things := range pl {
		pf := vh.PrjFiles(things)
		report.CSVReport(things, pf)

		semaphoreChan := make(chan struct{}, 20) // a blocking channel to keep concurrency under control
		defer close(semaphoreChan)
		wg := sync.WaitGroup{} // a wait group enables the main process a wait for goroutines to finish

		for k := range pf.Holdings {
			wg.Add(1)
			//log.Printf("About to run #%d in a goroutine\n", k)
			go func(k int) {
				semaphoreChan <- struct{}{}

				var n int64
				var l int
				// If the type is unknown, if it is a dir or starts with a . then skip it..
				if pf.Holdings[k].Type != "Exclude" && pf.Holdings[k].Type != "Directory" && !strings.HasPrefix(pf.Holdings[k].FileName, ".") {
					shaval := utils.SHAFile(pf.Holdings[k].Name)
					guid := xid.New()

					if pf.Holdings[k].Age > 2.00 {
						l = report.RDFGraph(guid.String(), pf.Holdings[k], shaval, &b) // this is for the "full" graph file written..   (do I still need this?)

						// TODO  load each metadata graph to minio..  then load those into Jena later
						var lb utils.Buffer
						_ = report.RDFGraph(guid.String(), pf.Holdings[k], shaval, &lb)

						jld, err := utils.NQToJSONLD(lb.String())
						if err != nil {
							log.Printf("Error converting NQ to JSON-LD: %v\n", err)
						}

						// b := bytes.NewBufferString(lb.String())  // when sending NQ, convert the string to a io reader bytes buffer string
						b := bytes.NewBuffer(jld) // if conversting lb to JSON-LD then that comes back as byte array, so make a new byte buffer

						contentType := "application/ld+json" // really Nq right now
						//n, err := mc.PutObject("doa-meta", objectName, b, int64(b.Len()), minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})
						n, err := mc.PutObject(fmt.Sprintf("%s-meta", bucketVal), guid.String(), b, int64(b.Len()), gominio.PutObjectOptions{ContentType: contentType})
						log.Printf("Loading metadata object: %d\n", n)
						if err != nil {
							log.Printf("Error loading metadata object to minio bucket %s : %s\n", bucketVal, err)
						}
					}

					// TODO uploading the metadata is easy..   just make a new buffer, get the triples, convert to JSON-LD loaded to a bucket

					if uploadVal && pf.Holdings[k].Age > 2.00 {
						n, err = minio.LoadToMinio(pf.Holdings[k].Name, bucketVal, pf.Holdings[k].FileName, pf.Holdings[k].Project, pf.Holdings[k].Type, pf.Holdings[k].FileExt, shaval, guid.String(), mc)
						log.Printf("Loading DO: %d\n", n)
						if err != nil {
							log.Printf("Error loading digital object to minio bucket %s : %s\n", bucketVal, err)
						}
					}
					//tr = append(tr, dg...)
					//	fmt.Printf("\nProject: %s\nType: %s\nName: %s\nRel: %s\nDir: %s\nFile: %s\nExt: %s \n",
					//		item.Project, item.Type, item.Name, item.RelativePath, item.ParentDir, item.FileName, item.FileExt)
				}

				log.Printf("RDF graph len: %d, Minio write len: %d File age: %f Routine %d\n", l, n, pf.Holdings[k].Age, k)
				wg.Done() // tell the wait group that we be done
				<-semaphoreChan
			}(k)
		}
		wg.Wait()
	}

	log.Println(b.Len())
	utils.WriteRDF(b.String())
	elapsed := time.Since(start)
	log.Printf("Walker took %s", elapsed)
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
		}
		*files = append(*files, path)
		return nil
	}
}
