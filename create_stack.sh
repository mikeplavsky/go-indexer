#!/bin/bash

ES_STACK_NUM=1
let ES_STACK_NUM=ES_STACK_NUM-1

for i in $(seq 0 $ES_STACK_NUM);
do 

let port=i+8080

docker create \
-m 15g \
-v /data/data$i:/data \
-p $port:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=3g \
--name=es$i --privileged --restart=always mikeplavsky/elasticsearch

done
