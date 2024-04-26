#!/bin/sh
PUSH=$1
DOCKER_REPO=spectrenetwork/spectred

set -e

tag=$(git log -n1 --format="%cs.%h")

cd $(dirname $(cd $(dirname $0); pwd))
docker build --pull -t $DOCKER_REPO:$tag -f docker/Dockerfile .
docker tag $DOCKER_REPO:$tag $DOCKER_REPO:latest
echo Tagged $DOCKER_REPO:latest

if [ "$PUSH" = "push" ]; then
  docker push $DOCKER_REPO:$tag
  docker push $DOCKER_REPO:latest
fi
