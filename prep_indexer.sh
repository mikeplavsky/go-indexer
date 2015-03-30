#!/bin/bash

while ! docker build --no-cache -t go_indexer .
do
	Sleep 1
done


