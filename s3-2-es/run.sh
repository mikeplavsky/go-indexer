#!/bin/sh

THROTTLING_LINES_PER_UPLOAD=100000

echo "Getting list of S3 files"
go-s3 -ls $S3_BUCKET/ > /tmp/all-items

echo "Generating bulk requests"
s3-2-es -in /tmp/all-items -out /tmp/all-requests -index $INDEX

rm -rf /tmp/x*
split -d -l $THROTTLING_LINES_PER_UPLOAD /tmp/all-requests

echo "Posting data"
for f in /tmp/x*; 
do 

curl -XPOST es:9200/_bulk --data-binary @/data/$f > /data/all-requests-$f.log ; 

done 

