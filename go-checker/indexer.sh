#!/bin/bash

ES_QUEUE="i-794345ae_dl"

go-checker

if [ $? -ne 0 ] 
then
	echo $S3_PATH | go-send s
fi


