// +build ignore

// Copyright 2015-2017 Piprate Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/piprate/json-gold/ld"
)

func main() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	// add the processing mode explicitly if you need JSON-LD 1.1 features
	options.ProcessingMode = ld.JsonLd_1_1

	triples := `
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/description> "Data set description" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/keywords> "DSDP, OPD, IODP, JanusAgeDatapoint" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/license> "https://creativecommons.org/publicdomain/zero/1.0/" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/name> "101_628A_JanusAgeDatapoint_CGcYexOw.csv" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/url> "http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://schema.org/Dataset> <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/distribution> _:b0 <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/publisher> _:b1 <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://schema.org/spatialCoverage> _:b2 <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://opencoredata.org/voc/janus/v1/Leg> "101" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://opencoredata.org/voc/janus/v1/Site> "628" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://opencoredata.org/voc/janus/v1/Hole> "A" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
<http://opencoredata.org/id/dataset/da39147d-deda-44ac-879d-684491a110fe> <http://opencoredata.org/voc/janus/v1/Measurement> "JanusAgeDatapoint" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b0 <http://schema.org/contentUrl> "http://opencoredata.org/api/v1/documents/download/204_1244B_JanusThermalConductivity_rkfjjNYV.csv" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b0 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> "http://schema.org/DataDownload" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b1 <http://schema.org/description> "NSF funded International Ocean Discovery Program operated by JRSO" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b1 <http://schema.org/name> "International Ocean Discovery Program" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b1 <http://schema.org/url> "http://iodp.org" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b1 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> "http://schema.org/Organization" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b2 <http://schema.org/geo> "b3" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> "http://schema.org/Place" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b3 <http://schema.org/latitude> "44.59" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b3 <http://schema.org/longitude> "-125.12" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
_:b3 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> "http://schema.org/GeoCoordinates" <http://opencoredata.org/objectgraph/id/da39147d-deda-44ac-879d-684491a110fe> .
	`

	doc, err := proc.FromRDF(triples, options)
	if err != nil {
		panic(err)
	}

	// ld.PrintDocument("JSON-LD output", doc)

	context := map[string]interface{}{
		"@context": map[string]interface{}{
			"dc":     "http://purl.org/dc/elements/1.1/",
			"schema": "http://schema.org/",
			"jrso":   "http://opencoredata.org/voc/janus/v1/",
		},
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		panic(err)

	}

	ld.PrintDocument("JSON-LD compation succeeded", compactedDoc)
}
