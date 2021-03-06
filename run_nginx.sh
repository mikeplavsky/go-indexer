docker rm -f nginx

docker create  \
-v $(pwd)/ui:/data/www \
-v $(pwd)/nginx.conf:/etc/nginx/nginx.conf:ro \
-v $(pwd)/.htpasswd:/etc/nginx/.htpasswd:ro \
-p 443:443 \
--name=nginx --net=host --restart=always mikeplavsky/docker-nginx

docker start nginx
