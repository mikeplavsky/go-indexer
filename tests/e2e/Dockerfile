FROM dockerfile/java:oracle-java7

RUN \
  curl -o fitnesse-standalone.jar -L http://fitnesse.org/fitnesse-standalone.jar?responder=releaseDownload&release=20150226 && \
  sudo apt-get update && \
  sudo apt-get install -y python2.7 && \
  sudo apt-get install -y build-essential python-dev libxml2-dev libxslt-dev  && \
  sudo apt-get install ca-certificates && \
  WAFERSLIM_URL=http://pypi.python.org/packages/source/w/waferslim  && \
  WAFERSLIM_URL=http://pypi.python.org/packages/source/w/waferslim && \
  WAFERSLIM_ZIP=waferslim-1.0.2-py2.6.zip && \
  WAFERSLIM_DIR=waferslim-1.0.2 && \
  DIST_PACKAGES_DIR=/usr/local/lib/python2.7/dist-packages && \

  WAFERSLIM_ZIP_URL=$WAFERSLIM_URL/$WAFERSLIM_ZIP && \

  wget $WAFERSLIM_ZIP_URL && \
  sudo apt-get install unzip -y && \
  unzip $WAFERSLIM_ZIP && \
  sudo rsync -r ./$WAFERSLIM_DIR/waferslim  $DIST_PACKAGES_DIR/ && \
  rm -rf $WAFERSLIM_ZIP && \
  rm -rf $WAFERSLIM_DIR 

EXPOSE 3680

ADD / /data

CMD ./trust_ca.sh && java -jar fitnesse-standalone.jar -p 3680
