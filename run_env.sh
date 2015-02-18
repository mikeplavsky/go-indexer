docker run -ti \
-e ES_INDEX=test0  \
-e ES_PORT=8080  \
-e ES_QUEUE=$ES_QUEUE \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=$ES_QUEUE \
-e ES_INDEXER=./indexer.sh \
--net=host --rm --name=go-env --rm go_indexer bash  
