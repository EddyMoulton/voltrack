#!/bin/bash

script_directory=$(dirname ${BASH_SOURCE[0]})
src_directory=$(dirname "$script_directory")

source $script_directory/variables.sh

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
  docker build -f "$src_directory/dockerfile.production" -t $REGISTRY/$IMAGE:latest $src_directory
else
  echo "Building development image"
  docker build -f "$src_directory/dockerfile" -t $REGISTRY/$IMAGE:latest $src_directory
fi
