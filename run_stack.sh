docker run \
-d \
-v /data/data0:/data \
-p 8080:9200 \
-e ES_CLUSTER_NAME=es0 \
-e ES_HEAP_SIZE=1g \
--name=es0 --restart=always mikeplavsky/elasticsearch

