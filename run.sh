#!/bin/sh

while ! curl -XPUT localhost:$ES_PORT/$ES_INDEX --data-binary @index.json 
do
	sleep 1
done

./go-indexer
