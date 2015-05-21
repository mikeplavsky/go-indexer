#!/bin/bash

PUBLIC_IP=$(curl http://169.254.169.254/latest/meta-data/public-ipv4)

sed "s/\$PUBLIC_IP/$PUBLIC_IP/g" route53.json > route53_i.json

sed    "s/\$AWS_STACK_NAME/$AWS_STACK_NAME/g" nginxT.conf > nginx.conf
sed -i "s/\$AWS_STACK_NAME/$AWS_STACK_NAME/g" route53_i.json

aws route53 change-resource-record-sets \
--hosted-zone-id Z2XLKK5Y4BC9Z1 \
--change-batch file://route53_i.json --region us-east-1
