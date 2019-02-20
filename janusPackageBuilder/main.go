package main

import (
	"fmt"

	"opencoredata.org/ocdGarden/janusPackageBuilder/internal/build"
	_ "opencoredata.org/ocdGarden/janusPackageBuilder/internal/fetch"
	_ "opencoredata.org/ocdGarden/janusPackageBuilder/internal/utils"
)

func main() {
	fmt.Println("....")

	// mc := utils.MinioConnection("192.168.2.131:9000", "AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

	// read the sitemap from Gleaner pkg files
	// fetch.ReadSitemap(mc)

	build.Pkg()

	// fetch each URL and extract JSON-LD document
	// bucket: meta  UUID named object for schema.org (UUID from URL)
	// bucket: data  UUID named object for data (UUID from URL)

	// build the packages in memory filesystem
	// bucket: package UUID named package
}
