#!/bin/bash

function help() {
  echo "Usage: [-v verbose]"
  echo "-v: verbose"
}

# read flags.
while getopts vh option
do
  case $option in
  v)  VERBOSE="true";;
  *)  help
      exit 1
  esac
done

if [ "$VERBOSE" = "true" ]; then
  ab -n 100 -c 4 -S -q 'localhost:32345/h1?a=1&b=3'
else
  ab -n 100 -c 4 -S -q 'localhost:32345/h1?a=1&b=3' | grep 'Time taken for tests:'
fi
