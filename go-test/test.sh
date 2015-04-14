#!/bin/bash

set -e
set -o pipefail

echo begin

unzip -p "MAgE_20150128_173131.zip" | go-convert | split -d -l 100000

echo end
