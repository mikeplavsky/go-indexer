FROM mikeplavsky/docker-golang

RUN mkdir /root/.ssh/
ADD id_rsa /root/.ssh/id_rsa

RUN chmod 600 /root/.ssh/id_rsa && \
    touch /root/.ssh/known_hosts && \
    apt-get update -y && \	
    apt-get install -y openssh-client && \	
    ssh-keyscan bitbucket.org >> /root/.ssh/known_hosts	

RUN git clone git@bitbucket.org:maplpro/go-convert.git src/go-convert && \
    git clone git@bitbucket.org:maplpro/go-indexer.git src/go-indexer && \
    go get -d go-indexer && \
    go install go-indexer && \	
    go install go-convert	

ADD index.json /go/bin/
ADD run.sh /go/bin/

WORKDIR /go/bin
CMD ./run.sh 
