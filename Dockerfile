FROM mikeplavsky/docker-golang

RUN apt-get update -y && \
    apt-get install unzip -y 

ENV UPDATED 09.03.2015.1

COPY . /go/src/go-indexer/
WORKDIR /go/src/go-indexer

RUN go get -t ./... && \
    go install ./... && \
    go test ./...	

CMD ./run.sh 

