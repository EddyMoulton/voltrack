#!/bin/bash

script_directory=$(dirname ${BASH_SOURCE[0]})
source $script_directory/variables.sh
src_directory=$(dirname "$script_directory")

# ensure we're up to date
#git pull

# bump version
current_version=$(cat $src_directory/VERSION)
echo $current_version | awk -F. -v OFS=. 'NF==1{print ++$NF}; NF>1{if(length($NF+1)>length($NF))$(NF-1)++; $NF=sprintf("%0*d", length($NF), ($NF+1)%(10^length($NF))); print}' >$src_directory/VERSION
version=$(cat $src_directory/VERSION)
echo "version: $version"

# run build
$script_directory/build.sh --production

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
