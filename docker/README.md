# Docker basics


## 1 - Why docker 

- call?

## 2 - Simple dockerfile

1. Create a dockerfile to build counter.go file
2. deploy the docker image on local computer and connect it to localhost:8080
3. try out localhost:8080/count, localhost:8080/increase, localhost:8080/decrease

## 3 - Multi stage dockerfile

1. Create dockerfile for counter.go, but use multi stage build, the final build should only have the binary and nothing else.

## 4 - Push to gitlab cr

1. Push the earlier created image to gitlab container registry https://gitlab.com/moonblade/jw-basics/container_registry

