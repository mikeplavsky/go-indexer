
docker run \
-d \
-m 16g \
-v /data/data0:/data \
-p 8080:9200 \
-e ES_CLUSTER_NAME=es0 \
-e ES_HEAP_SIZE=15g \
--name=es0 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data1:/data \
-p 8081:9200 \
-e ES_CLUSTER_NAME=es1 \
-e ES_HEAP_SIZE=15g \
--name=es1 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data2:/data \
-p 8082:9200 \
-e ES_CLUSTER_NAME=es2 \
-e ES_HEAP_SIZE=15g \
--name=es2 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data3:/data \
-p 8083:9200 \
-e ES_CLUSTER_NAME=es3 \
-e ES_HEAP_SIZE=15g \
--name=es3 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data4:/data \
-p 8084:9200 \
-e ES_CLUSTER_NAME=es4 \
-e ES_HEAP_SIZE=15g \
--name=es4 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data5:/data \
-p 8085:9200 \
-e ES_CLUSTER_NAME=es5 \
-e ES_HEAP_SIZE=15g \
--name=es5 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data6:/data \
-p 8086:9200 \
-e ES_CLUSTER_NAME=es6 \
-e ES_HEAP_SIZE=15g \
--name=es6 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data7:/data \
-p 8087:9200 \
-e ES_CLUSTER_NAME=es7 \
-e ES_HEAP_SIZE=15g \
--name=es7 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data8:/data \
-p 8088:9200 \
-e ES_CLUSTER_NAME=es8 \
-e ES_HEAP_SIZE=15g \
--name=es8 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data9:/data \
-p 8089:9200 \
-e ES_CLUSTER_NAME=es9 \
-e ES_HEAP_SIZE=15g \
--name=es9 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data10:/data \
-p 8090:9200 \
-e ES_CLUSTER_NAME=es10 \
-e ES_HEAP_SIZE=15g \
--name=es10 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data11:/data \
-p 8091:9200 \
-e ES_CLUSTER_NAME=es11 \
-e ES_HEAP_SIZE=15g \
--name=es11 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data12:/data \
-p 8092:9200 \
-e ES_CLUSTER_NAME=es12 \
-e ES_HEAP_SIZE=15g \
--name=es12 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data13:/data \
-p 8093:9200 \
-e ES_CLUSTER_NAME=es13 \
-e ES_HEAP_SIZE=15g \
--name=es13 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data14:/data \
-p 8094:9200 \
-e ES_CLUSTER_NAME=es14 \
-e ES_HEAP_SIZE=15g \
--name=es14 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data15:/data \
-p 8095:9200 \
-e ES_CLUSTER_NAME=es15 \
-e ES_HEAP_SIZE=15g \
--name=es15 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data16:/data \
-p 8096:9200 \
-e ES_CLUSTER_NAME=es16 \
-e ES_HEAP_SIZE=15g \
--name=es16 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data17:/data \
-p 8097:9200 \
-e ES_CLUSTER_NAME=es17 \
-e ES_HEAP_SIZE=15g \
--name=es17 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data18:/data \
-p 8098:9200 \
-e ES_CLUSTER_NAME=es18 \
-e ES_HEAP_SIZE=15g \
--name=es18 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data19:/data \
-p 8099:9200 \
-e ES_CLUSTER_NAME=es19 \
-e ES_HEAP_SIZE=15g \
--name=es19 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data20:/data \
-p 8100:9200 \
-e ES_CLUSTER_NAME=es20 \
-e ES_HEAP_SIZE=15g \
--name=es20 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data21:/data \
-p 8101:9200 \
-e ES_CLUSTER_NAME=es21 \
-e ES_HEAP_SIZE=15g \
--name=es21 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data22:/data \
-p 8102:9200 \
-e ES_CLUSTER_NAME=es22 \
-e ES_HEAP_SIZE=15g \
--name=es22 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data23:/data \
-p 8103:9200 \
-e ES_CLUSTER_NAME=es23 \
-e ES_HEAP_SIZE=15g \
--name=es23 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data24:/data \
-p 8104:9200 \
-e ES_CLUSTER_NAME=es24 \
-e ES_HEAP_SIZE=15g \
--name=es24 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data25:/data \
-p 8105:9200 \
-e ES_CLUSTER_NAME=es25 \
-e ES_HEAP_SIZE=15g \
--name=es25 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data26:/data \
-p 8106:9200 \
-e ES_CLUSTER_NAME=es26 \
-e ES_HEAP_SIZE=15g \
--name=es26 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data27:/data \
-p 8107:9200 \
-e ES_CLUSTER_NAME=es27 \
-e ES_HEAP_SIZE=15g \
--name=es27 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data28:/data \
-p 8108:9200 \
-e ES_CLUSTER_NAME=es28 \
-e ES_HEAP_SIZE=15g \
--name=es28 --restart=always mikeplavsky/elasticsearch

docker run \
-d \
-m 16g \
-v /data/data29:/data \
-p 8109:9200 \
-e ES_CLUSTER_NAME=es29 \
-e ES_HEAP_SIZE=15g \
--name=es29 --restart=always mikeplavsky/elasticsearch

