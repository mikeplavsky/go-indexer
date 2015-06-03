docker rm -f fit

docker run -d \
--net=host \
-v $(pwd)/FitNesseRoot/LogMan:/FitNesseRoot/LogMan \
--name fit mikeplavsky/docker-waferslim
