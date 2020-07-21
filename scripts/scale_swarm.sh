#!/bin/bash
set -e

function help() {
  echo "Usage: [-s service] [-r replicas]"
  echo "-s: service id"
  echo "-r: replicas number"
}

# read flags.
while getopts s:r:h option
do
  case $option in
  s)  SVC=OPTARG;;
  r)  REP=$OPTARG;;
  *)  help
      exit 1
  esac
done

if [ "$SVC" = "1" ]; then
  docker service update --replicas $REP demo_h1-1
elif [ "$SVC" = "2" ]; then
  docker service update --replicas $REP demo_h1-2 
elif [ "$SVC" = "3" ]; then
  docker service update --replicas $REP demo_h1-3 
else
  docker service update --replicas $REP demo_h1-1 
  docker service update --replicas $REP demo_h1-2 
  docker service update --replicas $REP demo_h1-3 
fi
