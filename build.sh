#!/bin/bash

source variables.sh

production_mode=0

while test $# -gt 0; do
  case "$1" in
  --production)
    production_mode=1
    ;;
  *)
    echo "argument $1"
    ;;
  esac
  shift
done

if [ $production_mode -eq 1 ]; then
  echo "Building production image"
  docker build -f dockerfile.production -t $REGISTRY/$IMAGE:latest .
else
  echo "Building development image"
  docker build -t $REGISTRY/$IMAGE:latest .
fi
