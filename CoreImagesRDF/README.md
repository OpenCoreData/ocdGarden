# About

A program that takes the old RDF I worked up for the core images Chris Jenkins worked on and updates it.

 
## Some regex notes

Some commands I used to conver the old RDF

http://brg.ldeo.columbia.edu/imagestrips/    ->  http://opencoredata.org/api/v1/image/download/
<http://data.oceandrilling.org/core/1/   ->    <http://opencoredata.org/voc/janus/1/
<http://data.oceandrilling.org/core/imagestrips/1/#  ->  <http://opencoredata.org/voc/janus/1/
<http://data.oceandrilling.org/coreimagery/  ->  <http://opencoredata.org/doc/imageset/


sed -i 's/foo/bar/g' file.ttl

