### IndexBuilder

#### About
A bleve index builder...  
Calls into Mongo in this case to do the indexing


#### SPARQL call for param info

http://opencoredata.org/blazegraph/namespace/opencore/sparql?query=


SELECT  ?uri ?name ?type ?column ?desc WHERE {  
?uri <http://example.org/rdf/type> <http://opencoredata.org/id/voc/janus/v1/JanusQuerySet> .   
?uri     <http://opencoredata.org/id/voc/janus/v1/struct_name> "janusvcdimage" .  
?uri   <http://opencoredata.org/id/voc/janus/v1/go_struct_name> ?name .
?uri  <http://opencoredata.org/id/voc/janus/v1/go_struct_type> ?type .  
?uri    <http://opencoredata.org/id/voc/janus/v1/column_id> ?column  .
?uri    <http://opencoredata.org/id/voc/janus/v1/JanusMeasurement> ?jmes .  
?jmes  <http://opencoredata.org/id/voc/janus/v1/json_descript>  ?desc  
}
ORDER By (xsd:integer(?column))