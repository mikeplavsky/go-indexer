while ! curl -XGET localhost:8080
do
        sleep 1
done

curl -XDELETE localhost:8080/s3data

docker rm s3-2-es

docker run \
-e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
-e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
--name s3-2-es \
--net=host \
-w /go/src/go-indexer/s3-2-es \
go_indexer ./run.sh
