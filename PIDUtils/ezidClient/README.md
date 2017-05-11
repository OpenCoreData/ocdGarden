# EZID client testing for Open Core DataCite


### About
This is a simple Go based client to submit a single XML datacite payload to EZID.  It's a bit of 
test code being used to develop a process to register resources in Open Core Data with 
EZID.  


### Testing notes from API docs
EZID provides two namespaces (or "shoulders") for testing purposes: ark:/99999/fk4 for ARK identifiers and doi:10.5072/FK2 for DOI identifiers. Identifiers in these namespaces are termed "test identifiers." They are ordinary long-term identifiers in almost all respects, including resolvability, except that EZID deletes them after 2 weeks.

Test DOI identifiers resolve through the universal DOI resolver (http://doi.org/), but do not appear in any of DataCite's other systems. Test DOI identifiers registered with Crossref appear only in Crossref's test server (http://test.crossref.org/), and are prefixed there with 10.15697. For example, test identifier doi:10.5072/FK2TEST will appear as doi:10.15697/10.5072/FK2TEST in Crossref.

All user accounts are permitted to create test identifiers. EZID also provides an "apitest" account that is permitted to create only test identifiers. Contact UC3 for the password for this account.


### V4 Issue
```
bash-3.2$ go run main.go                                                                                                                                                                
URL:> https://ezid.cdlib.org/shoulder/doi:10.5072/FK2                                                                                                                                   
response Status: 400 BAD REQUEST                                                                                                                                                        
response Headers: map[Content-Type:[text/plain; charset=UTF-8] Date:[Mon, 03 Oct 2016 02:00:11 GMT] Server:[Apache/2.2.17 (Unix) mod_ssl/2.2.17 OpenSSL/1.0.1k-fips mod_wsgi/4.4.9 Pytho
n/2.7.6] Content-Length:[76] Vary:[Accept-Language,Cookie] Content-Language:[en]]                                                                                                       
response Body: error: bad request - element 'datacite': unsupported DataCite record version                                                                                             


bash-3.2$ go run main.go                                                                                                                                                                
URL:> https://ezid.cdlib.org/shoulder/doi:10.5072/FK2                                                                                                                                   
response Status: 201 CREATED                                                                                                                                                            
response Headers: map[Vary:[Accept-Language,Cookie] Content-Language:[en] Content-Type:[text/plain; charset=UTF-8] Date:[Mon, 03 Oct 2016 02:00:24 GMT] Server:[Apache/2.2.17 (Unix) mod
_ssl/2.2.17 OpenSSL/1.0.1k-fips mod_wsgi/4.4.9 Python/2.7.6] Content-Length:[55]]                                                                                                       
response Body: success: doi:10.5072/FK2N29V896 | ark:/b5072/fk2n29v896                    
```


```
<?xml version="1.0"?>
<resource xmlns="http://datacite.org/schema/kernel-3"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="...">
  <identifier identifierType="DOI">(:tba)</identifier>
  ...
</resource>
```

```
<?xml version="1.0" encoding="UTF-8"?>
<records>
  <record identifier="ark:/99999/fk4gt78tq">
    <element name="_created">1300812337</element>
    <element name="_export">yes</element>
    <element name="_owner">apitest</element>
    <element name="_ownergroup">apitest</element>
    <element name="_profile">erc</element>
    <element name="_status">public</element>
    <element name="_target">http://www.gutenberg.org/ebooks/7178</element>
    <element name="_updated">1300913550</element>
    <element name="erc.what">Remembrance of Things Past</element>
    <element name="erc.when">1922</element>
    <element name="erc.who">Proust, Marcel</element>
  </record>
  <record identifier="doi:10.5072/FK2S75905Q">
    <element name="_created">1421276359</element>
    <element name="_datacenter">CDL.CDL</element>
    <element name="_export">yes</element>
    <element name="_owner">apitest</element>
    <element name="_ownergroup">apitest</element>
    <element name="_profile">datacite</element>
    <element name="_shadowedby">ark:/b5072/fk2s75905q</element>
    <element name="_status">public</element>
    <element name="_target">http://www.gutenberg.org/ebooks/26014</element>
    <element name="_updated">1421276359</element>
    <element name="datacite">
      <resource xmlns="http://datacite.org/schema/kernel-3">
        <identifier identifierType="DOI">10.5072/FK2S75905Q</identifier>
        <creators>
          <creator>
            <creatorName>Montagu Browne</creatorName>
          </creator>
        </creators>
        <titles>
          <title>Practical Taxidermy</title>
        </titles>
        <publisher>Charles Scribner's Sons</publisher>
        <publicationYear>1884</publicationYear>
        <resourceType resourceTypeGeneral="Text"/>
      </resource>
    </element>
  </record>
</records>
```
