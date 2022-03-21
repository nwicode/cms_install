#!/usr/bin/env bash

echo "################################## Run nginx"
export DOLLAR='$'
envsubst < ./src/nginx.conf > /etc/nginx/nginx.conf
nginx -g "daemon off;"