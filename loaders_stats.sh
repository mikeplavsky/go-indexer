#!/bin/bash

ES_STACK_NUM=$(cat /proc/cpuinfo | grep processor | wc -l)
let ES_STACK_NUM=ES_STACK_NUM-1

all=""

for i in $(seq 0 $ES_STACK_NUM);
do
all+=" loader$i"
done

docker stats $all
