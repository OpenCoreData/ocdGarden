## CSDCO Directory sevices


### Running
In the working directory make the output directories

* output
* output/kv
* output/pakcages  (only needed with -package)

The log file and report will be in `output`.  The kv directory will hold the KV
store, and the packages directory holds the generated packages if the -package flag is present.  

```
./vaultwalker -dir /media/fils/seagate/CSDdata -resetkv -index -report

IMPORTANT:  There is a bug issue, where the directory path must not end in /
This will be fixed, but is still present
```

On first run you may see an error like 
```
deleting error: bucket not found
```
This is fine and can be ignored..  I'll handle this edge case later.


### Options

```
./vaultwalker -help
Usage of ./vaultwalker:
  -dir string
        directory to index (default ".")
  -graph
        a bool for graph building
  -package
        a bool for package building
  -report
        a bool for report build
  -resetkv
        a bool for reseting the KV store
```

### Index and review sheet

Check the report.xlsx file in the output directory.  See if this will provide the sort of
review you need. 


### graph

Notes and steps
1. [ ] one
1. [ ] two
1. [ ] three
1. [ ] four



### packaging
