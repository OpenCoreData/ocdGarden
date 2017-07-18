### Prov Notes

#### Intro
Trying to learn what prov data might look like for OCD.



#### Refs

* https://gist.github.com/fils/3d337cd3768342646376206b7d5ac873 


```

@prefix csvw: <http://www.w3.org/ns/csvw#> .
@prefix prov: <http://www.w3.org/ns/prov#> .
@prefix xsd: <http://www.w3.org/2001/XMLSchema#> .

<http://opencoredata.org/doc/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb/prov>   # is a /prov extension really a good way to define a URI for this?
    a prov:Attribution ;
    prov:agent <http://doi.org/10.17616/R37936> ;  # does this need to be made a node with schema type DOI noted?
    prov:hadRole "Publisher" ;        # is this a litterial ?
    prov:wasAssociatedWith  <https://github.com/OpenCoreData> ;  # just a URL..  need a DOI for this git project?
    prov:qualifiedUsage  <http://opencoredata.org/doc/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb/prov#qu1> ;
    prov:qualifiedUsage  <http://opencoredata.org/doc/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb/prov#qu2> .

<http://opencoredata.org/doc/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb/prov#qu1>
       a prov:Usage ;
       prov:entity <http://opencoredata.org/api/v1/documents/download/199_1215A_JanusVcdImage_JcAruSDk.csv> ;
       prov:hadRole csvw:csvEncodedTabularData .

<http://opencoredata.org/doc/dataset/045deec9-94b2-445a-8fd2-43dbe90841fb/prov#qu2>
       a prov:Usage ;
       prov:entity <http://opencoredata.org/api/v1/documents/download/045deec9-94b2-445a-8fd2-43dbe90841fb/CSVW> ;
       prov:hadRole csvw:tabularMetadata .

```
