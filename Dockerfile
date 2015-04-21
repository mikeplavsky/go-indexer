FROM mikeplavsky/docker-golang

ENV UPDATED 15.04.2015.1

COPY . /go/src/go-indexer/
WORKDIR /go/src/go-indexer

RUN go get -t ./...

RUN cd /go/src/github.com/stretchr/testify && \
    git reset --hard e4ec8152c15fc46bd5056ce65997a07c7d415325

RUN go test ./...
RUN go install ./...

CMD ./run.sh 

