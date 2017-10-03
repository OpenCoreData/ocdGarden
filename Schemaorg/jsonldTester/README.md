### JSON-LD Tester

A simple bit of code that takes a URL, extracts the JSON-LD and then parses
it to see if it is well formed.  I would love to be able to see if it is 
valid against a vocabulary like schema.org or others.  That is more complex
of course. 

##### Testing

Some early test with this code in single call (non-concurrent mode) resulted 
in a time for 18171 (OCD on XSEDE) calls of: 

2017/10/03 10:58:36 P418 indexer took 54m38.789842318s

Later, a concurrent versiion set to allow X simultaneous calls resulted in a
time for 18171 (OCD on XSEDE) calls of:




##### http 2.0

Should try and do this via http2.0 and see what looks like...