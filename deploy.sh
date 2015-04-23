#!/bin/bash

git archive -o go-indexer.zip HEAD
aws s3 cp go-indexer.zip s3://logmanagement/
rm go-indexer.zip
