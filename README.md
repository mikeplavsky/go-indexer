# go-indexer
Docker container for indexing of zipped log files from AWS S3

### Choose instance type. 

Criteria:

* Number of events = Number of files * 70K
* Time to index = Number of events / Speed
* One log file ~ 70K events
* Index Size ~ Unziped Logs Size / 3

Resources needed:
- Memory: 6.5GB per CPU 
- 48 IOPS per CPU

Current speeds:

- 2.5K events per second per CPU

Current Speed per instance type:

- r3.2xlarge (8CPU, Memory: 60GB, Disk: 128GB): 
    - 20K events per second
    - 72M of events per hour
    - 1028 log files per hour

- r3.8xlarge (32CPU, Memory: 250GB, Disk: 512GB)
    - 80K events per second
    - 280M of events per hous
    - 4112 log files per hour

Steps to run:

- Install Docker https://docs.docker.com/installation/ubuntulinux/

    - For Ubuntu: curl -sSL https://get.docker.com/ubuntu/ | sudo sh 
    - http://askubuntu.com/questions/477551/how-can-i-use-docker-without-sudo
      - sudo groupadd docker
      - sudo gpasswd -a ${USER} docker
      - sudo service docker restart
      - newgrp docker

- Set Environment Variables, if IAM role is used - AWS_... variables are not needed

    - export AWS_ACCESS_KEY_ID=...
    - export AWS_SECRET_ACCESS_KEY=...
    
- Prepare Images

    - git clone git@github.com:GitQuest/go-indexer.git
    - cd go-indexer/
    - ./prep_env.sh

- Run
    
    - ./run_stack.sh
     
- Use App :)

    - Go to https://IP

- In Kibana there are two index patterns:
    
    - s3data This one is where S3 was synced 
    - test* This is where all logs are. 
