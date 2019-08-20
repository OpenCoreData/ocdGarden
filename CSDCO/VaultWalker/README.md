# Vaultwalker

## About
A new code base for the directory walk, filter and package needs of CSDCO


## Walker
Walk the vault and load the files

## Grapher
Take the sqlite file and convert to a graph for the associated resources
the facility defines  (look at stonesoup as guidance)

* Build the projects as objects (since they will have landing pages)
* Build the borehole features as objects (are they samples? Do them as sample 
landing pages)

Load these as objects (type data graph) into minio to then be loaded into the
triple store

Modify the existing ocdGarden/CSDCO/GraphBuilder code base?  It builds to a graph
but should be able to build data graph objects.   Or do I really need to follow that
pattern here?  It might be nice to still follow this path if that is what we want
in BCO-DMO.


## Commands

```
source ./secret/local.env
./cmd/walker/walker -d=/media/fils/T5/DataFiles/CSDCO/CSDCOdata -upload=true
```

```
go run cmd/walker/main.go -d=/media/fils/T5/DataFiles/CSDCO/CSDCOdata -upload=true
``

## Graph notes

Note: ObjectEngine is the place where the graphs are generated and loaded (tika and meta data loading)

```
curl -X POST --header "Content-Type:application/n-quads" -d @./output/objectGraph.nq http://192.168.2.132:3030/doa/data?graph=run1
```

```
SELECT distinct ?g ?s ?p ?o
WHERE {
   graph ?g{
       ?s ?p ?o 
    }
 }
limit 25
```
