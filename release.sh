#!/bin/bash

source variables.sh

# ensure we're up to date
#git pull

# bump version
docker run --rm -v "$PWD":/app treeder/bump patch
version=$(cat VERSION)
echo "version: $version"

# run build
./build.sh --production

# tag it
#git add -A
#git commit -m "version $version"
#git tag -a "$version" -m "version $version"
#git push
#git push --tags
docker tag $REGISTRY/$IMAGE:latest $REGISTRY/$IMAGE:$version

# push it
docker push $REGISTRY/$IMAGE:latest
docker push $REGISTRY/$IMAGE:$version
