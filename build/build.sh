#!/bin/bash

# Default options
target="all"
env="development"

# Handle arugments and loading common settings
/bin/bash common.sh

if [ "$target" = "all" ] || [ "$target" = "api" ]; then
  if [ "$env" = "development" ]; then
    echo "Building development API image"
    docker build -f "$script_directory/dockerfile.api.development" -t $REGISTRY/$IMAGE_API:latest $src_directory
  elif [ "$env" = "production" ]; then
    echo "Building production API image"
    docker build -f "$script_directory/dockerfile.api.production" -t $REGISTRY/$IMAGE_API:latest $src_directory
  fi
fi

if [ "$target" = "all" ] || [ "$target" = "web" ]; then
  if [ "$env" = "development" ]; then
    echo "Building development web app image"
    docker build -f "$script_directory/dockerfile.web.development" -t $REGISTRY/$IMAGE_WEB:latest $src_directory
  elif [ "$env" = "production" ]; then
    echo "Building production web app image"
    docker build -f "$script_directory/dockerfile.web.production" -t $REGISTRY/$IMAGE_WEB:latest $src_directory
  fi
fi
