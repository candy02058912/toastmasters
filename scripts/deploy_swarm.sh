#!/bin/bash
set -e
DIR=$(echo $(cd $(dirname ${BASH_SOURCE[0]})/.. && pwd))

docker build -t local/nginx -f ${DIR}/docker/nginx.Dockerfile ${DIR}/src/nginx
docker build -t local/h1 -f ${DIR}/docker/backend.Dockerfile ${DIR}/src/backend
docker build -t local/tester -f ${DIR}/docker/tester.Dockerfile ${DIR}/scripts

{ 
  docker swarm init 2> /dev/null 
} || { 
  echo 'swarm mode is already init.' 
}

docker stack deploy -c ${DIR}/docker-compose/docker-compose.yaml demo
