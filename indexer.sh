#!/bin/sh

#ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)-count
#unzip -p $ES_FILE | go-convert | wc -l | awk '{print "count/"$0}' | go-send s

unzip -p $ES_FILE | go-convert > /tmp/mage.json
curl -S -XPOST localhost:$ES_PORT/$ES_INDEX/_bulk --data-binary @/tmp/mage.json > /dev/null
