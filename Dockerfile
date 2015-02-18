FROM mikeplavsky/docker-golang

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

