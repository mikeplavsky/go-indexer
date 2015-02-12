FROM centos

RUN yum install nc -y
VOLUME ["/app"]

WORKDIR /app
CMD /app/run.sh
