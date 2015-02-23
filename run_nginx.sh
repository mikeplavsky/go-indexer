docker create  \
-v $(pwd)/ui:/data/www \
-v $(pwd)/nginx.conf:/etc/nginx/nginx.conf:ro \
-p 443:443 \
--name=nginx --net=host --restart=always mikeplavsky/docker-nginx

docker start nginx
