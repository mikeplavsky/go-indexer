docker pull mikeplavsky/elasticsearch
docker pull mikeplavsky/kibana
docker pull mikeplavsky/docker-golang
docker pull mikeplavsky/docker-nginx

docker build --no-cache -t go_indexer .


