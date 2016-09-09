package main

import (
	// "opencoredata.org/ocdGarden/skosBuilder/parameters"
	"opencoredata.org/ocdGarden/skosBuilder/janusQuerries"
	"opencoredata.org/ocdGarden/skosBuilder/parametersv2"
)

func main() {
	// parameters.Parameters()
	parametersv2.Parametersv2()
	janusQuerries.JanusQuerries()
}
