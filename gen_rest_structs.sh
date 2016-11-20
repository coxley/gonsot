#!/bin/bash

# Generate structs by calling local NSoT instance and using gojson

a=(
"localhost:8990/api/sites/"
"localhost:8990/api/attributes/"
"localhost:8990/api/devices/"
"localhost:8990/api/interfaces/"
"localhost:8990/api/circuits/"
"localhost:8990/api/networks/"
"localhost:8990/api/users/"
"localhost:8990/api/changes/"
"localhost:8990/api/values/"
)

cat <(for i in ${a[@]}; do
    r=${i:(19):(-2)}  # Resource name
    curl -s -L -H "X-NSoT-Email: coxley@blah.com" ${i} \
        | jq '.[0]' \
        | gojson -name=${r[@]^} -subStruct -pkg=rest \
        | sed 's/^package.*//g'
done) > rest/gen_types.go
