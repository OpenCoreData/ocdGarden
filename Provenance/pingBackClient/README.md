# pingBackClient

### About
This is a simple Go based client built simply to exercise from prov pingback capabilities a site may expose.  It will attempt to get, inspect and respond to a resource based on the prov-aw pingback pattern.


### Flow via postman

##### Step 1
Let's say we GET a resoruce like ``` http://127.0.0.1:9900/rdf/graph/void.ttl ```

Within the HEADER we would expect something like:

```
Link →<http://opencoredata.org/id/rdf/graph/void.ttl/provenance>; rel="http://www.w3.org/ns/prov#has_provenance"
Link →<http://opencoredata.org/rdf/rdf/graph/void.ttl/pingback>; rel="http://www.w3.org/ns/prov#pingbck"
```

##### Step 2
So, we would be able to do 2 things.

1) Look at the existing provenance with a GET to  http://127.0.0.1:9900/id/rdf/graph/void.ttl/provenance

2) Send in our new provenance for the record with a POST to http://127.0.0.1:9900/rdf/graph/void.ttl/pingback


For the latter we would expect a 204 (no content) return and accept that as a successful request/response event.   

