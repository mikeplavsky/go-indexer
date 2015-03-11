#!/bin/sh

ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)-error

unzip -p $ES_FILE | go-convert > /tmp/mage.json

while ! curl -S -XPOST localhost:$ES_PORT/$ES_INDEX/_bulk --data-binary @/tmp/mage.json > /dev/null
do 
echo "error/$ES_INDEX" | go-send s
done
