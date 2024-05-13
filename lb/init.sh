#!/bin/bash

# Assign default values to HOST_IP, RANKING_PORT, and SEARCH_PORT
export RANKING_PORT=8000
export SEARCH_PORT=9200

while getopts h:r:s: flag
do
    case "${flag}" in 
        h) HOST_IP=${OPTARG};;
        r) RANKING_PORT=${OPTARG};;
        s) SEARCH_PORT=${OPTARG};;
    esac 
done     


echo HOST_IP:$HOST_IP
echo RANKING_PORT:$RANKING_PORT
echo SEARCH_PORT:$SEARCH_PORT

envsubst < $CUR_DIR/lb/nginx.conf.template > $CUR_DIR/lb/nginx.conf