docker rm -f esspeed

docker run -d \
-e ES_INSTANCE_ID=$(curl http://169.254.169.254/latest/meta-data/instance-id) \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
--name esspeed \
--net=host \
-w /go/src/go-indexer/esspeed \
go_indexer esspeed
