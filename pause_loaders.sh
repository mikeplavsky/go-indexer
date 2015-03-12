#!/bin/bash

ES_STACK_NUM=$(cat /proc/cpuinfo | grep processor | wc -l)
let ES_STACK_NUM=ES_STACK_NUM-1

for i in $(seq 0 $ES_STACK_NUM);
do

docker pause loader$i

done
