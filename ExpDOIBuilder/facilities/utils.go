package facilities

import (
	"bufio"
	"log"
	"os"

	"github.com/knakk/rdf"
)

func ObjectValForPred(sparql map[string][]rdf.Term, predcol string, objectcol string, predicate string) string {
	pSlice := sparql[predcol]
	oSlice := sparql[objectcol]
	indexval := SliceIndex(len(pSlice), func(i int) bool { return pSlice[i].String() == predicate })
	// fmt.Println(indexval)
	// if -1 above return what?
	if indexval == -1 {
		return ""
	}
	return oSlice[indexval].String()
}

// SliceIndex return int location of item in slice  http://stackoverflow.com/questions/8307478/go-how-to-find-out-element-position-in-slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func writeFile(name string, xmldata string) {
	// Create the output file
	outFile, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	w := bufio.NewWriter(outFile)

	_, err = w.WriteString(xmldata)
	w.Flush()

	if err != nil {
		log.Fatal(err)
	}
}
