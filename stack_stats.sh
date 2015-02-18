#!/bin/bash

all=""

for i in $(seq 0 $ES_STACK_NUM);
do
all+=" es$i"
done

docker stats $all
