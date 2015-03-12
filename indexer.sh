#!/bin/sh

THROTTLING_LINES_PER_UPLOAD=100000
rm -rf ./x*

unzip -p $ES_FILE | go-convert > /tmp/mage.json
split -d -l $THROTTLING_LINES_PER_UPLOAD /tmp/mage.json

echo "Posting data"

for f in ./x*; 
do
	echo $f;
	curl -S -XPOST localhost:$ES_PORT/$ES_INDEX/_bulk --data-binary @$f > /dev/null;
done

