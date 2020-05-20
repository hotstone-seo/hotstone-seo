#!/bin/bash

TMP=$(mktemp)
OUT=$(./typicalw r 2> $TMP)
ERR=$(<$TMP)

echo "TMP:"
echo $TMP

echo "STDOUT:"
echo $OUT

echo "STDERR:"
echo $ERR

if [[ $ERR == *"connection refused"* ]]; then
  exit 0
else
  exit 1
fi
