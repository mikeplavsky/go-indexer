#!/bin/bash

ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)

ES_STACK_NUM=$(cat /proc/cpuinfo | grep processor | wc -l)
let ES_STACK_NUM=ES_STACK_NUM-1

for i in $(seq 0 $ES_STACK_NUM);
do

let port=i+8080

docker create \
-v $(pwd)/indexer.sh:/go/src/go-indexer/indexer.sh:ro \
-v $(pwd)/index.json:/go/src/go-indexer/index.json:ro \
-e ES_INDEX=test$i \
-e ES_PORT=$port \
-e ES_QUEUE=$ES_QUEUE \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader$i go_indexer 

done
