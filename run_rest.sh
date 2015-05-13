docker run -d \
-e ES_QUEUE=$(curl http://169.254.169.254/latest/meta-data/instance-id) \
--name rest-srv \
--net=host \
--restart=always go_indexer rest-srv
