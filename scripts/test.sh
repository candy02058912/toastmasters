#!/bin/bash

set -e
DIR=$(echo $(cd $(dirname ${BASH_SOURCE[0]})/.. && pwd))

{
    docker swarm leave --force
} || {
    echo '...'
}
docker swarm init
# docker-compose -f ${DIR}/docker-compose/docker-compose.yaml up -d
docker stack deploy -c ${DIR}/docker-compose/docker-compose.yaml demo
ab -n 1200 -c 4 -S -q -l 'localhost:32345/h1?a=1&b=3'

# ```
# Creating service demo_strawberry
# failed to create service demo_strawberry: Error response from daemon: network demo_backend not found
# ```