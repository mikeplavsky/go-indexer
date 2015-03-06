docker rm -f fit

docker build -t fit .

#docker create  \
#-p 3680:3680 \
#--name=fit fit 

docker run -d --net=host -p 3680:3680 -v $(pwd)/:/data fit
