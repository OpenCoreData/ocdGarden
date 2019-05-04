#!/bin/bash
#a wrapper script for invoking 
mc_cmd() {
        mc ls clear/csdco-graphs | awk '{print $5}'
}

for i in $(mc_cmd); do
    echo "$i"
    mc cat clear/csdco-graphs/$i | curl -X POST --header "Content-Type:application/n-triples" -d @- http://localhost:3030/doa/data?graph=test
done

