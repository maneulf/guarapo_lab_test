#! /bin/bash

echo "Generating certs for https server"

mkdir -p ../gitignore/certs #ruta donde se ponen cert and key

openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ../certs/key.pem -out ../certs/cert.pem

echo "Generated certs"