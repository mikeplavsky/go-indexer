#!/bin/sh

while ! nc -w 1 localhost $ES_PORT
do
   echo -n .
   sleep 1	
done

curl -XPUT localhost:$ES_PORT/$ES_INDEX --data-binary @index.json
./go-indexer
