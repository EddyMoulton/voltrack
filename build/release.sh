#!/bin/bash

script_directory=$(dirname $(realpath -s $0))
src_directory=$(dirname "$script_directory")

source $script_directory/common.sh

# Parse aguments with defaults: target = all, env = production
parse_args $@

if [ -z $target ]; then
  target="all"
fi

if [ -z $env ]; then
  env="production"
fi

if [ "$env" != "production" ]; then
  echo "May only release for production"
  exit 1
fi

if [[ -n $(git status -s) ]] && [[ -z $skipCommitCheck ]]; then
  echo "Uncommitted changes: Ensure all changes are committed before releasing"
  exit 1
fi

cd $src_directory

# Increment version
current_version=$(cat VERSION)
echo $current_version | awk -F. -v OFS=. 'NF==1{print ++$NF}; NF>1{if(length($NF+1)>length($NF))$(NF-1)++; $NF=sprintf("%0*d", length($NF), ($NF+1)%(10^length($NF))); print}' >VERSION
echo "version: $current_version"

# Run build
./build/build.sh --environment "$env" --target "$target"

# Tag in git
git add -A
git commit -m "version $current_version"
git tag -a "$current_version" -m "version $current_version"

# Push to remote - disabled due to permissions
git push
git push --tags

if [ "$target" = "all" ] || [ "$target" = "api" ]; then
  docker tag $REGISTRY/$IMAGE_API:latest $REGISTRY/$IMAGE_API:$current_version
  docker push $REGISTRY/$IMAGE_API:latest
  docker push $REGISTRY/$IMAGE_API:$current_version
fi

if [ "$target" = "all" ] || [ "$target" = "web" ]; then
  docker tag $REGISTRY/$IMAGE_WEB:latest $REGISTRY/$IMAGE_WEB:$current_version
  docker push $REGISTRY/$IMAGE_WEB:latest
  docker push $REGISTRY/$IMAGE_WEB:$current_version
fi

cd -
