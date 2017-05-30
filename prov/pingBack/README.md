# Prov-aq testing

### About
A simple package to play with some prov-aq

https://www.w3.org/TR/prov-aq/
https://www.w3.org/TR/prov-aq/#provenance-pingback

### Examples from above resources to consider 

We would get a resource like
```
C: GET http://acme.example.org/super-widget123 HTTP/1.1

S: 200 OK
S: Link: <http://acme.example.org/super-widget123/provenance>; 
         rel="http://www.w3.org/ns/prov#has_provenance"
S: Link: <http://acme.example.org/super-widget123/pingback>; 
         rel="http://www.w3.org/ns/prov#pingback"
 :
(super-widget123 resource data)
```

so we need to get some resources out there with these headers.


```
C: POST http://acme.example.org/super-widget123/pingback HTTP/1.1
C: Content-Type: text/uri-list
C:
C: http://coyote.example.org/contraption/provenance
C: http://coyote.example.org/another/provenance

S: 204 No Content
```

```
C: POST http://acme.example.org/super-widget123/pingback HTTP/1.1
C: Link: <http://coyote.example.org/extra/provenance>;
         rel="http://www.w3.org/ns/prov#has_provenance";
         anchor="http://acme.example.org/extra-widget"
C: Content-Type: text/uri-list
C:
C: http://coyote.example.org/contraption/provenance
C: http://coyote.example.org/another/provenance
C: http://coyote.example.org/extra/provenance

S: 204 No Content
```
