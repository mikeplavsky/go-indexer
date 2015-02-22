FROM mikeplavsky/docker-golang

RUN apt-get update -y && \
    apt-get install unzip -y && \
    apt-get install -y vim	 

ENV UPDATED 22.02.2015.3

COPY . /go/src/go-indexer/
WORKDIR /go/src/go-indexer

RUN go get ./... && \
    go install ./...

CMD ./run.sh 

