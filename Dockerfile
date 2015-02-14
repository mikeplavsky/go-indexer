FROM mikeplavsky/docker-golang

RUN mkdir /root/.ssh/
ADD id_rsa /root/.ssh/id_rsa

RUN chmod 600 /root/.ssh/id_rsa && \
    touch /root/.ssh/known_hosts && \
    apt-get update -y && \	
    apt-get install -y openssh-client && \	
    ssh-keyscan bitbucket.org >> /root/.ssh/known_hosts && \
    apt-get install unzip -y && \
    apt-get install -y vim		

ENV UPDATED=15.02.2015.0

RUN git clone git@bitbucket.org:maplpro/go-convert.git src/go-convert && \
    git clone git@bitbucket.org:maplpro/go-indexer.git src/go-indexer && \
    git clone git@bitbucket.org:maplpro/go-s3.git src/go-s3 && \
    git clone git@bitbucket.org:maplpro/go-send.git src/go-send && \
    go get -d go-indexer && \
    go get -d go-s3 && \
    go get -d go-send && \
    go install go-s3 && \	
    go install go-send && \	
    go install go-indexer && \	
    go install go-convert	

WORKDIR /go/src/go-indexer
CMD ./run.sh 
