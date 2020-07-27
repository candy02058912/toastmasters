#!/bin/bash
set -e

function help() {
  echo "Usage: [-s service] [-r replicas]"
  echo "-s: service id"
  echo "-r: replicas number"
}

# read flags.
SVC=()
REP=()
while getopts s:r:h option
do
  case $option in
  s)  SVC+=($OPTARG);;
  r)  REP+=($OPTARG);;
  *)  help
      exit 1
  esac
done

if [ ${#SVC[@]} = "0" ]; then
  echo "scale all services."
  if [ ${#REP[@]} = "1" ]; then
    docker service update --replicas $REP demo_plain 
    docker service update --replicas $REP demo_chocolate
    docker service update --replicas $REP demo_strawberry
  else
    echo "should only have one -r if not specify service id."
    exit 1
  fi
elif [ ${#SVC[@]} != ${#REP[@]} ]; then
  echo "numbers of service id and replica mismatch."
  exit 1
fi

NUM=${#SVC[@]}
for (( i=0; i<$NUM; i++ ))
do
  if [ "${SVC[i]}" == "1" ]; then
    docker service update --replicas ${REP[$i]} demo_plain
  elif [ ${SVC[$i]} == "2" ]; then
    docker service update --replicas ${REP[$i]} demo_chocolate
  elif [ ${SVC[$i]} == "3" ]; then
    docker service update --replicas ${REP[$i]} demo_strawberry
  fi
done
