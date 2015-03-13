#!/bin/sh

while ! curl localhost:$ES_PORT
do
	sleep 1
done

go-indexer
