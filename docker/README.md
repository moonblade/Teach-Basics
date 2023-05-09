# Docker basics


## 1 - Why docker 

- add infra to application

## 1.1 - Terminology

- Images and containers
- Client and daemon
- Registry

## 1.2 - Helloworld

1. Create a dockerfile that prints Hello world with bash

## 2 - Simple dockerfile

1. Create a dockerfile to build counter.go file
2. deploy the docker image on local computer and connect it to localhost:8080
3. try out localhost:8080/count, localhost:8080/increase, localhost:8080/decrease

## 2.1 - Notes: Networks

1. Different networks
2. cidr
3. hostname and ip and ports - dns
4. network interfaces

## 3 - Multi stage dockerfile

1. Create dockerfile for counter.go, but use multi stage build, the final build should only have the binary and nothing else.
2. Create it with a non default name of multi-stage-dockerfile

## 4 - Push to docker io

1. Push the earlier image to docker.io, with the image name counter, and tag as 1.0.0-scratch

## 5 - Multiple images and env variables

1. Create dockerfile for persistant-counter, it needs another container with mongod installed to be running in parallel and the second container ip to be given thorugh env variable MONGO_HOST
- Note: microservice architecture

## 6. Push to gitlab cr
1. Push the earlier created image to gitlab container registry https://gitlab.com/moonblade/jw-basics/container_registry
2. Push to local docker registry as well

## 7 - Docker compose

1. Do task 5, but use docker compose to accomplish the connections between different dockers

## 8 - Persistance

1. Was the data in task 5 persisted across restarts? if not redo it such that each time the instance starts it should start with the previous value of count last saved 

