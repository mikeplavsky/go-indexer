ES_STACK_NUM=$(cat /proc/cpuinfo | grep processor | wc -l)
ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id)

./prep_indexer.sh
./route53.sh

./create_queues.sh

./create_stack.sh
./create_loaders.sh

./start_stack.sh
./run_esspeed.sh

./run_s3.sh
./run_rest.sh
./run_kibana.sh
./run_nginx.sh

./start_loaders.sh

./run_sync.sh
./run_ipython.sh
