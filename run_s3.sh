docker run \
--name s3-2-es \
--net=host \
-w /go/src/go-indexer/s3-2-es \
go_indexer ./run.sh
