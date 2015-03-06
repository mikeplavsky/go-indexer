#!/bin/sh

echo -n | openssl s_client -connect localhost:443 | sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' > /tmp/server.cert
cp /tmp/server.cert /usr/local/share/ca-certificates/
sudo update-ca-certificates
