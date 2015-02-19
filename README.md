# go-indexer
Docker container for indexing of zipped log files from AWS S3

Steps to run:

- Install Docker https://docs.docker.com/installation/ubuntulinux/

    - For Ubuntu: curl -sSL https://get.docker.com/ubuntu/ | sudo sh 
    - http://askubuntu.com/questions/477551/how-can-i-use-docker-without-sudo
      - sudo groupadd docker
      - sudo gpasswd -a ${USER} docker
      - sudo service docker restart
      - newgrp docker

- Set Environment Variables 

    - export AWS_ACCESS_KEY_ID=...
    - export AWS_SECRET_ACCESS_KEY=...
    - export ES_QUEUE=...
    - export ES_STACK_NUM=...
    
- Copy your GitHub Key to .ssh
    - scp -i ~/.ssh/id_rsa ~/.ssh/id_rsa ubuntu@54.159.121.204:/home/ubuntu/.ssh/id_rsa
    - chmod 600 ~/.ssh/id_rsa

- Run Indexer
    
    - git clone git@github.com:GitQuest/go-indexer.git
    - cd go-indexer/
    - docker build -t go_indexer .
    - ./run_env.sh
    - go-s3 l dmp-log-analysis/D4755B98-A20C-42B1-A685-D42B5B326B52/folder/UnifiedMailSync_1 | go-send
    - ./create_stack.sh 
    - ./create_loaders.sh 
    - ./start_stack.sh
    - ./start_loaders.sh

- Rub Kibana 

    - ./run_kibana.sh
    - Go to https://IP
    
    
    


