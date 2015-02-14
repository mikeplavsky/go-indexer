
docker run \
-d \
-e ES_INDEX=test0 \
-e ES_PORT=8080 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader0 go_indexer 

docker run \
-d \
-e ES_INDEX=test1 \
-e ES_PORT=8081 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader1 go_indexer 

docker run \
-d \
-e ES_INDEX=test2 \
-e ES_PORT=8082 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader2 go_indexer 

docker run \
-d \
-e ES_INDEX=test3 \
-e ES_PORT=8083 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader3 go_indexer 

docker run \
-d \
-e ES_INDEX=test4 \
-e ES_PORT=8084 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader4 go_indexer 

docker run \
-d \
-e ES_INDEX=test5 \
-e ES_PORT=8085 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader5 go_indexer 

docker run \
-d \
-e ES_INDEX=test6 \
-e ES_PORT=8086 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader6 go_indexer 

docker run \
-d \
-e ES_INDEX=test7 \
-e ES_PORT=8087 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader7 go_indexer 

docker run \
-d \
-e ES_INDEX=test8 \
-e ES_PORT=8088 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader8 go_indexer 

docker run \
-d \
-e ES_INDEX=test9 \
-e ES_PORT=8089 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader9 go_indexer 

docker run \
-d \
-e ES_INDEX=test10 \
-e ES_PORT=8090 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader10 go_indexer 

docker run \
-d \
-e ES_INDEX=test11 \
-e ES_PORT=8091 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader11 go_indexer 

docker run \
-d \
-e ES_INDEX=test12 \
-e ES_PORT=8092 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader12 go_indexer 

docker run \
-d \
-e ES_INDEX=test13 \
-e ES_PORT=8093 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader13 go_indexer 

docker run \
-d \
-e ES_INDEX=test14 \
-e ES_PORT=8094 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader14 go_indexer 

