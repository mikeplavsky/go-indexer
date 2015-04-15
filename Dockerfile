FROM mikeplavsky/docker-golang

ENV UPDATED 15.04.2015.1

COPY . /go/src/go-indexer/
WORKDIR /go/src/go-indexer

RUN go get -t ./... && \
    go install ./... && \
    go test ./...	

CMD ./run.sh 

