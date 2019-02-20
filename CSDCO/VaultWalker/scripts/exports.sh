#!/bin/bash

while read p; do
      export $p
done <../secret/prod.env
