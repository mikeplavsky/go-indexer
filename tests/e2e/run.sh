docker rm -f fit

docker build -t fit .

docker run -d --net=host -p 3680:3680 -v $(pwd)/:/data -e COMMAND=$1 fit
