#!/bin/sh

export GIN_MODE=debug
export XFIV_APP_MODE="develop"

# file logs
# export XFIV_PATH_FILE_LOGS="../gitignore/logs/all.log"

# http and https
export ADDRESS_HTTP=":8080"
export ADDRESS_HTTPS=":4443"
export PATH_CERT_HTTPS="./certs/cert.pem"
export PATH_KEY_HTTPS="./certs/key.pem"
export READ_TIMEOUT="11"
export WRITE_TIMEOUT="11"
export PERSISTENCE_TYPE="sqlite"
#export PERSISTENCE_TYPE="inmemory"
