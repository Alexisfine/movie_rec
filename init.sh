#!/bin/bash

# Remove previous ENV file
ENV_FILE=".env"
rm -f $ENV_FILE

# Set up env args
export CUR_DIR=$(pwd)
export HOST_IP=$(ipconfig getifaddr en0)

# Create new nginx.conf file for the load balancer service 
./lb/init.sh 

# Start service
docker compose up -d
