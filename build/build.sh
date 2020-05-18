#!/bin/bash

# Default options
target="all"
env="development"

script_directory=$(dirname ${BASH_SOURCE[0]})
src_directory=$(dirname "$script_directory")

source $script_directory/variables.sh

# Handle arugments and loading common settings
/bin/bash $script_directory/common.sh

starting_dir="$PWD"
cd $src_directory

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

cd $starting_dir
