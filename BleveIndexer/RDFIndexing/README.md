# Index RDF with Bleve

# Notes


https://github.com/wallix/triplestore
https://github.com/deiu/rdf2go


# SPARQL

```
prefix bds: <http://www.bigdata.com/rdf/search#>
select DISTINCT ?name ?url ?progname ?description
where {
 
  { ?s <http://schema.org/name> ?name   .
  ?s <http://schema.org/url> ?url}
  
 UNION
  
 {?s <http://schema.org/programName> ?progname   . 
  ?s <http://schema.org/hostingOrganization> ?ho .
  ?ho <http://schema.org/url> ?url
 }
  
 UNION
  
 {?s <http://schema.org/description> ?description .
   ?s <http://schema.org/url> ?url
}

}
```