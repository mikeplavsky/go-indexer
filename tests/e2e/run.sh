docker rm -f fit

docker run -d \
--net=host \
-v $(pwd)/FitNesseRoot/LogMan:/FitNesseRoot/LogMan \
-v $(pwd)/fixtures:/fixtures \
-v $(pwd)/data:/data \
-v $(pwd)/aws:/aws \
--name fit mikeplavsky/docker-waferslim
