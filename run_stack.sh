
docker run \
-d \
-m 4g \
-v /data/data0:/data \
-p 8080:9200 \
-e ES_CLUSTER_NAME=es0 \
-e ES_HEAP_SIZE=3g \
--name=es0 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data1:/data \
-p 8081:9200 \
-e ES_CLUSTER_NAME=es1 \
-e ES_HEAP_SIZE=3g \
--name=es1 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data2:/data \
-p 8082:9200 \
-e ES_CLUSTER_NAME=es2 \
-e ES_HEAP_SIZE=3g \
--name=es2 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data3:/data \
-p 8083:9200 \
-e ES_CLUSTER_NAME=es3 \
-e ES_HEAP_SIZE=3g \
--name=es3 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data4:/data \
-p 8084:9200 \
-e ES_CLUSTER_NAME=es4 \
-e ES_HEAP_SIZE=3g \
--name=es4 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data5:/data \
-p 8085:9200 \
-e ES_CLUSTER_NAME=es5 \
-e ES_HEAP_SIZE=3g \
--name=es5 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data6:/data \
-p 8086:9200 \
-e ES_CLUSTER_NAME=es6 \
-e ES_HEAP_SIZE=3g \
--name=es6 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data7:/data \
-p 8087:9200 \
-e ES_CLUSTER_NAME=es7 \
-e ES_HEAP_SIZE=3g \
--name=es7 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data8:/data \
-p 8088:9200 \
-e ES_CLUSTER_NAME=es8 \
-e ES_HEAP_SIZE=3g \
--name=es8 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data9:/data \
-p 8089:9200 \
-e ES_CLUSTER_NAME=es9 \
-e ES_HEAP_SIZE=3g \
--name=es9 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data10:/data \
-p 8090:9200 \
-e ES_CLUSTER_NAME=es10 \
-e ES_HEAP_SIZE=3g \
--name=es10 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 4g \
-v /data/data11:/data \
-p 8091:9200 \
-e ES_CLUSTER_NAME=es11 \
-e ES_HEAP_SIZE=3g \
--name=es11 --restart=always mikeplavsky/elasticsearch

