package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/index"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/minio"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/report"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/internal/vault"
	"opencoredata.org/ocdGarden/CSDCO/VaultWalker/pkg/utils"
	//	minio "github.com/minio/minio-go"
)

var minioVal, portVal, accessVal, secretVal, bucketVal, dirVal string
var uploadVal, sslVal bool

func init() {
	akey := os.Getenv("MINIO_ACCESS_KEY")
	skey := os.Getenv("MINIO_SECRET_KEY")

	flag.StringVar(&minioVal, "address", "192.168.2.131", "FQDN for server")
	flag.StringVar(&portVal, "port", "9000", "Port for minio server, default 9000")
	flag.StringVar(&accessVal, "access", akey, "Access Key ID")
	flag.StringVar(&secretVal, "secret", skey, "Secret access key")
	flag.StringVar(&bucketVal, "bucket", "csdco", "The configuration bucket")
	flag.BoolVar(&sslVal, "ssl", false, "Use SSL boolean")
	flag.BoolVar(&uploadVal, "upload", false, "Upload files to object store")
	flag.StringVar(&dirVal, "d", "./test", "Directory to walk")
}

func main() {
	var files []string
	var va []vault.VaultItem

	flag.Parse()
	mc := minio.MinioConnection(minioVal, portVal, accessVal, secretVal)

	d := dirVal
	err := filepath.Walk(d, visit(&files))
	if err != nil {
		panic(err)
	}

	// make output directory if it doesn't exist
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.Mkdir("./output", os.ModePerm)
	}

	// get the various elements of the found files...
	for _, file := range files {
		v := index.PathInspection(d, file)
		va = append(va, v)
	}

	vh := vault.VaultHoldings{va}
	pl := vh.Prjs() // get unique project (dir) names

	fmt.Printf("{Projects: %s\n", pl)
	var b utils.Buffer

	for _, things := range pl {
		pf := vh.PrjFiles(things)
		report.CSVReport(things, pf)

		semaphoreChan := make(chan struct{}, 10) // a blocking channel to keep concurrency under control
		defer close(semaphoreChan)
		wg := sync.WaitGroup{} // a wait group enables the main process a wait for goroutines to finish

		for k := range pf.Holdings {

			wg.Add(1)
			log.Printf("About to run #%d in a goroutine\n", k)

			go func(k int) {
				semaphoreChan <- struct{}{}

				var n int64
				var l int
				// If the type is unknown, if it is a dir or starts with a . then skip it..
				if pf.Holdings[k].Type != "Unknown" && pf.Holdings[k].Type != "Directory" && !strings.HasPrefix(pf.Holdings[k].FileName, ".") {
					shaval := utils.SHAFile(pf.Holdings[k].Name)
					l = report.RDFGraph(pf.Holdings[k], shaval, &b) // need to expand the object graph
					if uploadVal {
						n, err = minio.LoadToMinio(pf.Holdings[k].Name, "csdco", pf.Holdings[k].FileName, pf.Holdings[k].Project, pf.Holdings[k].Type, pf.Holdings[k].FileExt, shaval, mc)
						if err != nil {
							log.Printf("Error loading to minio: %s\n", err)
						}
					}
					//tr = append(tr, dg...)
					//	fmt.Printf("\nProject: %s\nType: %s\nName: %s\nRel: %s\nDir: %s\nFile: %s\nExt: %s \n",
					//		item.Project, item.Type, item.Name, item.RelativePath, item.ParentDir, item.FileName, item.FileExt)
				}

				log.Printf("Buffer written len %d and minio write len %d by routine %d\n", l, n, k)
				wg.Done() // tell the wait group that we be done
				<-semaphoreChan
			}(k)
		}
		wg.Wait()
	}

	log.Println(b.Len())
	utils.WriteRDF(b.String())
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}
