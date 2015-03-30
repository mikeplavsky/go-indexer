ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)

docker rm -f queues

docker run \
-e ES_QUEUE=$ES_QUEUE \
--name=queues go_indexer go-send c
