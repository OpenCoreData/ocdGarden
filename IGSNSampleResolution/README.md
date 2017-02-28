## IGSN resolution 

### About
A simple program to resolve IGSN values to Open Core resources
 based on the sample IDs 


### Files

* odp_sample_id.tab:  File from LDEO with SESAR sample id and IGSN info
* output.txt:  results of matching these to URI's in data.oceandrilling.org 
* unknownSamples.txt:  Sample ID's from LDEO that are not present in the data.oceandrilling.org graph


### Notes
Samples in the graph are not unique by sample ID alone.  Rather, it looks like sample ID plus
repository is what makes the unique ID.  We need to look at those cases in the files and 
see what taking place more closely between the two data holdings.   