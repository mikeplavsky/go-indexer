docker run -ti \
-e ES_INDEX=test0  \
-e ES_PORT=8080  \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_INDEXER=./indexer.sh \
-v $(pwd):/go/src/go-indexer \
--net=host --rm --name=go-env --rm go_indexer bash  
