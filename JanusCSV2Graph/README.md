### JanusCSV2Graph

Some SPARQL call examples on the resulting graph

```
SELECT  DISTINCT *  
WHERE {
  ?uri <http://example.org/rdf/type> <http://opencoredata.org/id/voc/janus/v1/JanusQuerySet> .
  ?uri     <http://opencoredata.org/id/voc/janus/v1/struct_name> "janusageprofile" .
  ?uri ?p ?o .
}
```

Results with a 24 limit..  (will provide live SPARQL links when I get the graph deployed):

```
uri	p	o
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://example.org/rdf/type>	<http://opencoredata.org/id/voc/janus/v1/JanusQuerySet>
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://opencoredata.org/id/voc/janus/v1/column_id>	1
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://opencoredata.org/id/voc/janus/v1/column_name>	LEG
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://opencoredata.org/id/voc/janus/v1/go_struct_name>	Leg
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://opencoredata.org/id/voc/janus/v1/go_struct_type>	int64
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://opencoredata.org/id/voc/janus/v1/JanusMeasurement>	<http://opencoredata/id/resource/janus/measure/leg>
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://opencoredata.org/id/voc/janus/v1/struct_name>	janusageprofile
<http://opencoredata/id/resource/janus/janusageprofile/1>	<http://opencoredata.org/id/voc/janus/v1/table_name>	<http://opencoredata/id/resource/janus/query/ocd_age_profile>
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://example.org/rdf/type>	<http://opencoredata.org/id/voc/janus/v1/JanusQuerySet>
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://opencoredata.org/id/voc/janus/v1/column_id>	10
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://opencoredata.org/id/voc/janus/v1/column_name>	DATUM_ID
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://opencoredata.org/id/voc/janus/v1/go_struct_name>	Datum_id
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://opencoredata.org/id/voc/janus/v1/go_struct_type>	int64
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://opencoredata.org/id/voc/janus/v1/JanusMeasurement>	<http://opencoredata/id/resource/janus/measure/datum_id>
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://opencoredata.org/id/voc/janus/v1/struct_name>	janusageprofile
<http://opencoredata/id/resource/janus/janusageprofile/10>	<http://opencoredata.org/id/voc/janus/v1/table_name>	<http://opencoredata/id/resource/janus/query/ocd_age_profile>
<http://opencoredata/id/resource/janus/janusageprofile/11>	<http://example.org/rdf/type>	<http://opencoredata.org/id/voc/janus/v1/JanusQuerySet>
<http://opencoredata/id/resource/janus/janusageprofile/11>	<http://opencoredata.org/id/voc/janus/v1/column_name>	DATUM_TYPE
<http://opencoredata/id/resource/janus/janusageprofile/11>	<http://opencoredata.org/id/voc/janus/v1/go_struct_name>	Datum_type
<http://opencoredata/id/resource/janus/janusageprofile/11>	<http://opencoredata.org/id/voc/janus/v1/go_struct_type>	string
<http://opencoredata/id/resource/janus/janusageprofile/11>	<http://opencoredata.org/id/voc/janus/v1/JanusMeasurement>	<http://opencoredata/id/resource/janus/measure/datum_type>
<http://opencoredata/id/resource/janus/janusageprofile/11>	<http://opencoredata.org/id/voc/janus/v1/struct_name>	janusageprofile
<http://opencoredata/id/resource/janus/janusageprofile/11>	<http://opencoredata.org/id/voc/janus/v1/table_name>	<http://opencoredata/id/resource/janus/query/ocd_age_profile>
<http://opencoredata/id/resource/janus/janusageprofile/12>	<http://example.org/rdf/type>	<http://opencoredata.org/id/voc/janus/v1/JanusQuerySet>
```




Some other queries I am holding here for potential documentation.

```
prefix bds: <http://www.bigdata.com/rdf/search#>
select DISTINCT *
where {
   ?o bds:search "janustensorcore" .
   ?s ?p ?o .
   ?s ?pp ?oo .
}
```

```
prefix bds: <http://www.bigdata.com/rdf/search#>
select DISTINCT ?tn
where {
   ?o bds:search "janustensorcore" .
   ?s ?p ?o .
   ?s <http://opencoredata.org/id/voc/janus/v1/table_name>  ?tn
}
```

```
prefix bds: <http://www.bigdata.com/rdf/search#>
select ?s ?p ?o
where {
   ?o bds:search "janustensorcore" .
   ?s ?p ?o .
}
```

```
SELECT  *
WHERE {
  ?uri <http://example.org/rdf/type> <http://opencoredata.org/id/voc/janus/v1/JanusQuerySet> .
  ?uri     <http://opencoredata.org/id/voc/janus/v1/struct_name> "janusageprofile" .
  ?uri <http://opencoredata.org/id/voc/janus/v1/JanusMeasurement> ?jm .
  ?jm   ?p ?o
}
```

```
SELECT  *
WHERE {
  ?col  <http://opencoredata.org/id/voc/janus/v1/JanusMeasurement> <http://opencoredata/id/resource/janus/measure/datum_age> .
  ?col  <http://opencoredata.org/id/voc/janus/v1/table_name>  ?hdr .
  ?hdr ?p ?o  .   
}
```