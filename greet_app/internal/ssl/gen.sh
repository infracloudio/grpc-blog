#!/usr/bin/env bash

# Output files
# ca.key : Certificate Authority private key file [this should not be shared in real-life]
# ca.crt : Certificate Authority trust certificate (this should be shared with users in real-life)
# server.key : Server private key, password protected (this should not be shared)
# server.csr : Server certificate signing request (this should be shared with CA owner)
# server.crt : Sever certificate signed by the CA (this would be sent back by the CA owner) - keep on server
# server.pem : Conversion of server.key into format gRPC likes (this should not be shared)

# Summary
# Private files : ca.key, server.key, server.pem, server.crt
# "share" files : ca.crt (needed by client), server.csr (needed by the CA)

# Changes these CN's to match your hosts in your environment if needed
SERVER_CN=localhost

# step 1: generate certificate authority + trust certificate (ca.crt)
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj "//CN=${SERVER_CN}" -addext "subjectAltName=DNS:${SERVER_CN}"

# Step 2: generate the server private key (server.key)
openssl genrsa -passout pass:1111 -des3 -out server.key 4096

# Step 3: get a certificate signing request from the CA (server.csr)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "//CN=${SERVER_CN}" -addext "subjectAltName=DNS:${SERVER_CN}"

# Step 4: sign the certificate with the CA we created it is called self-signing (server.crt)
openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

# Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem