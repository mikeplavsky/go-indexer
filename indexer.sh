#!/bin/sh

unzip -p $ES_FILE | ./go-convert
curl -S -XPOST localhost:$ES_PORT/$ES_INDEX/_bulk --data-binary @/tmp/mage.json > /dev/null"
