#!/bin/bash

./build-service.sh

cp service-app ./docker/service-app

docker build -t registery-service:latest -f ./docker/Dockerfile ./docker/
