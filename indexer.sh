#!/bin/sh
set -e

curl -XPUT localhost:$ES_PORT/$ES_INDEX --data-binary @index.json 

THROTTLING_LINES_PER_UPLOAD=100000
rm -rf ./x*

unzip -p $ES_FILE | go-convert | split -d -l $THROTTLING_LINES_PER_UPLOAD

echo "Posting data"

for f in ./x*; 
do
	echo $f;
	curl -S -XPOST localhost:$ES_PORT/$ES_INDEX/_bulk --data-binary @$f > /dev/null;
done

