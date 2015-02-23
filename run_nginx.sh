docker create  \
-v $(pwd)/nginx.conf:/etc/nginx/nginx.conf:ro \
--name=nginx --net=host --restart=always mikeplavsky/docker-nginx

docker start nginx
