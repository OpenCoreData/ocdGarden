## SPARQL to DOT

## About
This was inspired by https://blog.semicolonsoftware.de/extending-ghost-with-graphviz/ and the 
idea of running a simple Docker container for graphviz that I could feed dot into and then 
view the results.  This might be nice in some of the SPARQL queries to give a visual representation 
of the results.


## REFS

* https://blog.semicolonsoftware.de/extending-ghost-with-graphviz/
* http://www.webgraphviz.com/
* https://github.com/awalterschulze/gographviz 

### Notes

```
Fils:SparqlToDot dfils$ fdp unique.dot -Tsvg -o image.svg

(dot|neato|twopi|circo|fdp|sfdp|patchwork|osage)


```

### SPARQL call used in the code

```
PREFIX schemaorg: <http://schema.org/>
SELECT DISTINCT ?repository  ?name ?memberOfName ?endpoint_url ?endpoint_description ?endpoint_method
WHERE {
  ?repository rdf:type <http://schema.org/Organization>   .
  ?repository schemaorg:name ?name   .
  ?repository schemaorg:memberOf ?mo  .
  ?mo   schemaorg:programName ?memberOfName   .
    ?repository schemaorg:potentialAction [ schemaorg:target ?action ] .
    ?action schemaorg:urlTemplate ?endpoint_url .
    ?action schemaorg:description ?endpoint_description .
    ?action schemaorg:httpMethod ?endpoint_method .

}
ORDER BY ?name
```