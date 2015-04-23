#!/bin/bash


InstanceId=$(curl http://169.254.169.254/latest/meta-data/instance-id)
ES_QUEUE=$InstanceId"_dl"

go-checker

if [ $? -ne 0 ] 
then
	echo $S3_PATH | go-send s
fi


