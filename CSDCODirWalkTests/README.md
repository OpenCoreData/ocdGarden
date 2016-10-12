# CSDCO Directory Walk testing

## STATUS Cultivated
This garden item has been cultivated.. its results are now in ocdFX

## About
Need to use the testing described below to limit the discovery of files to correctly index the 
CSDCO holdings and avoid file not of interest


## Simple patterns

```
in an sub-folder look for files with
PID-metadata
Dtube Label_PID
any file with SRF
any .cml
any .car

subdir Images
.jpg .jpeg  .tif .bmp

subdir Images/rgb
.csv

[Project]/Geotek Data/whole-core data/
contains _MSCL

[Project]/Geotek Data/high-resolution MS data/
_XYZ or _HRMS

[Project]/ICD  or below
.pdf

```


## Notes

```
These projects will have Geotek Data in our database, so the script can ignore that folder in these projects:

CLDP
CPCP
GLAD1
GLAD2
GLAD4
GLAD5
GLAD6
GLAD7
GLAD9
GLAD11
HOTSPOT
HSPDP
MEXI
ODP
OGDP
PLJ
TDP
TERUEL


The most important pathways in the Vault are:

[Project]/
(or subfolders)
any file with name = "[ProjectID]-metadata” will have two sheets; the first has hole-level metadata, keyed by HoleID; the second has section-level metadata, keyed by SectionID.
any file with name = "metadata format Dtube Label_[ProjectID]” will have section-level metadata, keyed by SectionID, with mostly duplicative information (compared to the above "[ProjectID]-metadata” file) but some additional fields (e.g. Diameter)
any file with a name that includes “SRF” has subsample metadata, keyed by SectionID. 
any .cml file is a Corelyzer session file (a few KB) that allows rapid inputs to Corelyzer.
any .car file is a Corelyzer archive file (hundred of MB to several GB) that includes all images in a single file for direct import.

[Project]/Images/
(or subfolders)
any .jpg is an image of a section, whose SectionID = the filename. 
earlier files will exclude the -W or -A at the end, which indicate ‘working’ or ‘archive’ half of the core. Recent files include these.  
any extraneous suffixes indicate special camera settings to capture some or all of the core more effectively (e.g. _lighter, _lighter2, _darker, _F4, _F8, _focus2, _focushigh, _focuslow, etc)
any .tif or .bmp is the uncompressed version of the jpeg. 
in earlier projects, only .tif or .bmp exist but not .jpeg versions.

[Project]/Images/rgb/
any .csv file contains the RGB data from a 5mm strip down the center of the section image whose name is indicated by the .csv filename.


[Project]/Geotek Data/whole-core data/
any file with “_MSCL” at the end of the filename is the whole-core multisensor logger data. 
SectionIDs are indicated in each row of the file
Usually, data from multiple boreholes is combined into a single file.


[Project]/Geotek Data/high-resolution MS data/
any file with “_XYZ” or “_HRMS” at the end of the filename is the split-core multisensor logger data. 
SectionIDs are indicated in each row of the file
Usually, data from multiple boreholes is combined into a single file.


[Project]/ICD/
any .pdf at this level or in subdirectories is a lithologic description of a section, whose SectionID = the filename. 
```

