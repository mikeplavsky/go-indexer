#!/bin/bash

for i in $(seq 0 $ES_STACK_NUM);
do

docker start loader$i

done
