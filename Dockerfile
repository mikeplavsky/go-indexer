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

ENV UPDATED=17.02.2015.4

RUN git clone git@bitbucket.org:maplpro/go-convert.git src/go-convert && \
    git clone git@bitbucket.org:maplpro/go-indexer.git src/go-indexer && \
    git clone git@bitbucket.org:maplpro/go-s3.git src/go-s3 && \
    git clone git@bitbucket.org:maplpro/go-send.git src/go-send && \
    git clone git@bitbucket.org:maplpro/esspeed.git src/esspeed && \
    git clone git@bitbucket.org:maplpro/repeater.git src/repeater && \
    go get -d go-indexer && \
    go get -d go-s3 && \
    go get -d go-send && \
    go get -d esspeed && \
    go get -d repeater && \
    go install go-s3 && \	
    go install repeater && \	
    go install esspeed && \	
    go install go-send && \	
    go install go-indexer && \	
    go install go-convert	

WORKDIR /go/src/go-indexer
CMD ./run.sh 
