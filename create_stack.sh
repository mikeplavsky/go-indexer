
docker create \
-m 15g \
-v /data/data0:/data \
-p 8080:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es0 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data1:/data \
-p 8081:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es1 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data2:/data \
-p 8082:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es2 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data3:/data \
-p 8083:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es3 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data4:/data \
-p 8084:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es4 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data5:/data \
-p 8085:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es5 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data6:/data \
-p 8086:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es6 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data7:/data \
-p 8087:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es7 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data8:/data \
-p 8088:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es8 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data9:/data \
-p 8089:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es9 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data10:/data \
-p 8090:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es10 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data11:/data \
-p 8091:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es11 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data12:/data \
-p 8092:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es12 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data13:/data \
-p 8093:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es13 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data14:/data \
-p 8094:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es14 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data15:/data \
-p 8095:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es15 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data16:/data \
-p 8096:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es16 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data17:/data \
-p 8097:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es17 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data18:/data \
-p 8098:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es18 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data19:/data \
-p 8099:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es19 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data20:/data \
-p 8100:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es20 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data21:/data \
-p 8101:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es21 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data22:/data \
-p 8102:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es22 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data23:/data \
-p 8103:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es23 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data24:/data \
-p 8104:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es24 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data25:/data \
-p 8105:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es25 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data26:/data \
-p 8106:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es26 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data27:/data \
-p 8107:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es27 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data28:/data \
-p 8108:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es28 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data29:/data \
-p 8109:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es29 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data30:/data \
-p 8110:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es30 --privileged --restart=always mikeplavsky/elasticsearch

docker create \
-m 15g \
-v /data/data31:/data \
-p 8111:9200 \
-e ES_CLUSTER_NAME=es \
-e ES_HEAP_SIZE=6g \
--name=es31 --privileged --restart=always mikeplavsky/elasticsearch

