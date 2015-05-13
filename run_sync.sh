docker rm -f go-sync

ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)

docker run -d \
-e ES_QUEUE=$ES_QUEUE"_sns" \
--name go-sync \
--net=host \
-w /go/src/go-indexer/go-sync \
--restart=always go_indexer ./run.sh 
