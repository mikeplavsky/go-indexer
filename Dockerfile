FROM mikeplavsky/docker-golang

ENV UPDATED 23.04.2015.1

COPY . /go/src/go-indexer/
WORKDIR /go/src/go-indexer

RUN godep restore
RUN go test ./...
RUN go install ./...

CMD ./run.sh 

