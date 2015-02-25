while ! curl -XGET localhost:8080
do
        sleep 1
done

curl -XDELETE localhost:8080/s3data

docker rm s3-2-es

docker run \
--name s3-2-es \
--net=host \
-w /go/src/go-indexer/s3-2-es \
go_indexer ./run.sh
