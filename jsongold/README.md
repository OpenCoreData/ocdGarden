# Playing with JSON-goLD package

## notes

When I convert my JSON-LD to RDF I was getting blank nodes.  To resolve this I give each section
an ID based on the resource URI with a #extention on them.  This then allows me to make SPARQL
queries.  So something like:

```
select DISTINCT *
where {
  {
    <http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28> ?p ?o .
  }
  UNION
  {
   <http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28#distribution> ?p ?o .
  }
   UNION
  {
    <http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28#place> ?p ?o .
  }
  UNION
  {
    <http://opencoredata.org/id/dataset/e8fb758e-22ba-499d-92fb-8d653febcf28#geocoordinates> ?p ?o .
  }
 }
 ```

 will pull back all the triples and allow the JSON-LD json format to be rebuilt.  

Having the results above coming back in SPARQL JSON format allows a rather easy process in Go to
place the results in array of structs that can be then be walked with the given query URI's and
then rebuild the nquads text.  This could be done via a text/template or using the Go RDF library
to build the triple (which would trap errors) and then feed these results into JSON-goLD to
build out the JSON text body.

While this might seem like a lot of work to do it's a validating path that both ensures our data
is valid and allow for error trapping along the way which would be harder with templates and raw
text.
