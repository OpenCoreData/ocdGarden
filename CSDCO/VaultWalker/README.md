# Vaultwalker

## About
A new code base for the directory walk, filter and package needs of CSDCO

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
