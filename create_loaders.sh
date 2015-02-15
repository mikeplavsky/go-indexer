
docker create \
-e ES_INDEX=test0 \
-e ES_PORT=8080 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader0 go_indexer 

docker create \
-e ES_INDEX=test1 \
-e ES_PORT=8081 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader1 go_indexer 

docker create \
-e ES_INDEX=test2 \
-e ES_PORT=8082 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader2 go_indexer 

docker create \
-e ES_INDEX=test3 \
-e ES_PORT=8083 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader3 go_indexer 

docker create \
-e ES_INDEX=test4 \
-e ES_PORT=8084 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader4 go_indexer 

docker create \
-e ES_INDEX=test5 \
-e ES_PORT=8085 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader5 go_indexer 

docker create \
-e ES_INDEX=test6 \
-e ES_PORT=8086 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader6 go_indexer 

docker create \
-e ES_INDEX=test7 \
-e ES_PORT=8087 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader7 go_indexer 

docker create \
-e ES_INDEX=test8 \
-e ES_PORT=8088 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader8 go_indexer 

docker create \
-e ES_INDEX=test9 \
-e ES_PORT=8089 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader9 go_indexer 

docker create \
-e ES_INDEX=test10 \
-e ES_PORT=8090 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader10 go_indexer 

docker create \
-e ES_INDEX=test11 \
-e ES_PORT=8091 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader11 go_indexer 

docker create \
-e ES_INDEX=test12 \
-e ES_PORT=8092 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader12 go_indexer 

docker create \
-e ES_INDEX=test13 \
-e ES_PORT=8093 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader13 go_indexer 

docker create \
-e ES_INDEX=test14 \
-e ES_PORT=8094 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader14 go_indexer 

docker create \
-e ES_INDEX=test15 \
-e ES_PORT=8095 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader15 go_indexer 

docker create \
-e ES_INDEX=test16 \
-e ES_PORT=8096 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader16 go_indexer 

docker create \
-e ES_INDEX=test17 \
-e ES_PORT=8097 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader17 go_indexer 

docker create \
-e ES_INDEX=test18 \
-e ES_PORT=8098 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader18 go_indexer 

docker create \
-e ES_INDEX=test19 \
-e ES_PORT=8099 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader19 go_indexer 

docker create \
-e ES_INDEX=test20 \
-e ES_PORT=8100 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader20 go_indexer 

docker create \
-e ES_INDEX=test21 \
-e ES_PORT=8101 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader21 go_indexer 

docker create \
-e ES_INDEX=test22 \
-e ES_PORT=8102 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader22 go_indexer 

docker create \
-e ES_INDEX=test23 \
-e ES_PORT=8103 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader23 go_indexer 

docker create \
-e ES_INDEX=test24 \
-e ES_PORT=8104 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader24 go_indexer 

docker create \
-e ES_INDEX=test25 \
-e ES_PORT=8105 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader25 go_indexer 

docker create \
-e ES_INDEX=test26 \
-e ES_PORT=8106 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader26 go_indexer 

docker create \
-e ES_INDEX=test27 \
-e ES_PORT=8107 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader27 go_indexer 

docker create \
-e ES_INDEX=test28 \
-e ES_PORT=8108 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader28 go_indexer 

docker create \
-e ES_INDEX=test29 \
-e ES_PORT=8109 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader29 go_indexer 

docker create \
-e ES_INDEX=test30 \
-e ES_PORT=8110 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader30 go_indexer 

docker create \
-e ES_INDEX=test31 \
-e ES_PORT=8111 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader31 go_indexer 

docker create \
-e ES_INDEX=test32 \
-e ES_PORT=8112 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader32 go_indexer 

docker create \
-e ES_INDEX=test33 \
-e ES_PORT=8113 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader33 go_indexer 

docker create \
-e ES_INDEX=test34 \
-e ES_PORT=8114 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader34 go_indexer 

docker create \
-e ES_INDEX=test35 \
-e ES_PORT=8115 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader35 go_indexer 

docker create \
-e ES_INDEX=test36 \
-e ES_PORT=8116 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader36 go_indexer 

docker create \
-e ES_INDEX=test37 \
-e ES_PORT=8117 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader37 go_indexer 

docker create \
-e ES_INDEX=test38 \
-e ES_PORT=8118 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader38 go_indexer 

docker create \
-e ES_INDEX=test39 \
-e ES_PORT=8119 \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
-e ES_QUEUE=lm-test \
-e ES_INDEXER=./indexer.sh \
--net=host --name=loader39 go_indexer 

