
docker run \
-d \
-v $(pwd):/app \
-e ES_INDEX=test0 \
-e ES_PORT=8080 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader0 go_indexer 

