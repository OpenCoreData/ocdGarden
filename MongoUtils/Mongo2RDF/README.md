# Mongo 2 RDF

### About
This is a fresh start on extracting data from Mongo into RDF.  It duplicates
some functionality found elsewhere in the garden and in the ocdBulk repo.  However, 
it's an attempt to do it both better and in a more organized manner.

### Notes

Items of interest
* abstracts -> csdco
* expedire -> expeditions
* expedire -> features
* expedire -> featuresABSGeoJSON
    * superset of featuresGeoJSON
* test -> csdco
* test -> csvmeta
* test -> jsonld
* test -> schemaorg
* test -> uniqueids


Existing RDF graphs 
* IODPPeople.ttl.gz
* JRSO_cruises_gl.ttl.gz
* JRSO_deployments_gl.ttl.gz
* JRSO_holes_gl.ttl.gz
* JanusDataTypes.nt.gz
* NGDC-DSDP.ttl.gz
* agedSections.ttl.gz
* chronosAgeModels.ttl.gz
* codices.nt.gz
* csdcoProjects.nt.gz
* geoLinkVoid.ttl.gz
* janusAmpNew.ttl.gz
* seas.ttl.gz