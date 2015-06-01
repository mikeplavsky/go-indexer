#!/bin/sh

THROTTLING_LINES_PER_UPLOAD=100000
ES_INDEX=s3data
S3_BUCKET=dmp-log-analysis/$AWS_CUSTOMER

while :
	do

	echo "Generating bulk requests"
	go-sync | grep S3_BUCKET | s3-2-es > /tmp/all-requests

	rm -rf ./x*
	split -d -l $THROTTLING_LINES_PER_UPLOAD /tmp/all-requests

	echo "Posting data"
	for f in ./x*; 
	do 
		echo $f;
		curl -XPOST localhost:8080/$ES_INDEX/_bulk --data-binary @$f > /dev/null;
	done 
done
