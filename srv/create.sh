#!/bin/bash

docker stop gotty
docker rm  gotty

docker build -t gotty .

docker run \
--mount type=bind,source="$(pwd)"/credentials.json,target=/credentials.json,readonly \
--detach \
--name=gotty \
--publish 8080:8080 \
--cap-add=sys_nice \
gotty
