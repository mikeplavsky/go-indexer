#/bin/sh

FROM=$1
TO=$2
CUSTOMER=$3

if [ -z "$FROM" ] || [ -z "$TO" ] || [ -z "$CUSTOMER" ] ; then 
	echo "usage: from to customer";
	echo "example: 2010-02-03 2014-04-05 CONTOSO";
	exit 1; 
fi

echo "pushing data to queue..."

export PATH=$PATH:$GOPATH/bin
s3-query -from=$FROM -to=$TO | go-send
