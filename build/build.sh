#!/bin/bash

script_directory=$(dirname ${BASH_SOURCE[0]})
src_directory=$(dirname "$script_directory")

source $script_directory/common.sh

# Parse aguments with defaults: target = all, env = development
parse_args $@

if [ -z $target ]; then
  target="all"
fi

if [ -z $env ]; then
  env="development"
fi

cd $src_directory

echo $target

if [ "$target" = "all" ] || [ "$target" = "api" ]; then
  if [ "$env" = "development" ]; then
    echo "Building development API image"
    docker build -f "build/dockerfile.api.development" -t $REGISTRY/$IMAGE_API:latest .
  elif [ "$env" = "production" ]; then
    echo "Building production API image"
    docker build -f "build/dockerfile.api.production" -t $REGISTRY/$IMAGE_API:latest .
  fi
fi

if [ "$target" = "all" ] || [ "$target" = "web" ]; then
  if [ "$env" = "development" ]; then
    echo "Building development web app image"
    docker build -f "build/dockerfile.web.development" -t $REGISTRY/$IMAGE_WEB:latest .
  elif [ "$env" = "production" ]; then
    echo "Building production web app image"
    docker build -f "build/dockerfile.web.production" -t $REGISTRY/$IMAGE_WEB:latest .
  fi
fi

cd -
