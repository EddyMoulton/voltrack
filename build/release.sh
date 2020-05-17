#!/bin/bash

# Default options
target="all"
env="production"

# Handle arugments and loading common settings
/bin/bash common.sh

if [ "$env" != "production" ]; then
  echo "May only release for production"
  exit 1
fi

# Increment version
current_version=$(cat $src_directory/VERSION)
echo $current_version | awk -F. -v OFS=. 'NF==1{print ++$NF}; NF>1{if(length($NF+1)>length($NF))$(NF-1)++; $NF=sprintf("%0*d", length($NF), ($NF+1)%(10^length($NF))); print}' >$src_directory/VERSION
version=$(cat $src_directory/VERSION)
echo "version: $version"

# Run build
$script_directory/build.sh --environment "$env" --target "$target"

# Tag in git
git add -A
git commit -m "version $version"
git tag -a "$version" -m "version $version"

# Push to remote - disabled due to permissions
#git push
#git push --tags

if [ "$target" = "all" ] || [ "$target" = "api" ]; then
  docker tag $REGISTRY/$IMAGE_API:latest $REGISTRY/$IMAGE_API:$version
  docker push $REGISTRY/$IMAGE_API:latest
  docker push $REGISTRY/$IMAGE_API:$version
fi

if [ "$target" = "all" ] || [ "$target" = "web" ]; then
  docker tag $REGISTRY/$IMAGE_WEB:latest $REGISTRY/$IMAGE_WEB:$version
  docker push $REGISTRY/$IMAGE_WEB:latest
  docker push $REGISTRY/$IMAGE_WEB:$version
fi
