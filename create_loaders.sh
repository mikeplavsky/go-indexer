#!/bin/bash

for i in $(seq 0 $ES_STACK_NUM);
do

let port=i+8080

docker create \
-e ES_INDEX=test$i \
-e ES_PORT=$port \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=$ES_QUEUE \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader$i go_indexer 

done
