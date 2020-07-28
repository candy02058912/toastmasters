#!/bin/bash

function help() {
  echo "Usage: [-t testcase, 1 or 2]"
  echo "-t: testcase, default to 1"
}

# read flags.
while getopts ht: option
do
  case $option in
  t)  TEST=$OPTARG;;
  *)  help
      exit 1
  esac
done

if [ "$TEST" = "1" ]; then
  echo "TEST 1: Low concurrency test"
  ab -n 50 -c 9 -S -q -l 'localhost:32345/h1'
elif [ "$TEST" = "2" ]; then
  echo "TEST 2: High concurrency test"
  ab -n 50 -c 50 -S -q -l 'localhost:32345/h1'
else 
  echo "TEST 1: Low concurrency test"
  ab -n 50 -c 9 -S -q -l 'localhost:32345/h1'
  echo "TEST 2: High concurrency test"
  ab -n 50 -c 50 -S -q -l 'localhost:32345/h1'
fi
