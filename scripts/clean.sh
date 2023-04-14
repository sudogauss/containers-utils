#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "You must provide only image name"
    exit 1
fi

docker rm $(docker stop $(docker ps -a -q --filter ancestor=$1 --format="{{.ID}}"))
docker rmi $1:latest