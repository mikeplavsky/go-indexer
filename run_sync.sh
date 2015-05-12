docker rm -f go-sync

ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)

docker run \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=$ES_QUEUE"_sns"
--name go-sync \
--net=host \
--restart=always go_indexer go-sync
