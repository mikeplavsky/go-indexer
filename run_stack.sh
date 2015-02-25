ES_STACK_NUM=$(cat /proc/cpuinfo | grep processor | wc -l)
ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)

docker run -ti \
-e ES_QUEUE=$ES_QUEUE \
--rm go_indexer go-send c

./create_stack.sh
./create_loaders.sh

./start_stack.sh
./start_loaders.sh

while ! curl -XGET localhost:8080
do
        sleep 1
done

./run_s3.sh
./run_rest.sh
./run_kibana.sh
./run_nginx.sh
