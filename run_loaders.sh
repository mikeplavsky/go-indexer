
docker run \
-d \
-v $(pwd):/app \
-e ES_INDEX=test0 \
-e ES_PORT=8080 \
--net=host --name=loader0 go-indexer 

