FROM mikeplavsky/docker-golang

RUN apt-get update -y && \
    apt-get install unzip -y && \
    apt-get install -y vim	 

ENV UPDATED=17.02.2015.7

COPY . /go/src/go-indexer/

COPY go-s3 /go/src/go-indexer/go-s3
COPY go-send /go/src/go-indexer/go-send
COPY repeater /go/src/go-indexer/repeater
COPY esspeed /go/src/go-indexer/esspeed

WORKDIR /go/src/go-indexer

RUN go get ./... && \
    go install ./...

CMD ./run.sh 

