FROM mikeplavsky/docker-golang

RUN apt-get update -y && \
    apt-get install unzip -y && \
    apt-get install -y vim	 

ENV UPDATED 21.02.2015.2

COPY . /go/src/go-indexer/

COPY go-s3 /go/src/go-indexer/go-s3
COPY go-send /go/src/go-indexer/go-send
COPY esspeed /go/src/go-indexer/esspeed
COPY s3-2-es /go/src/go-indexer/s3-2-es

WORKDIR /go/src/go-indexer

RUN go get ./... && \
    go install ./...

CMD ./run.sh 

