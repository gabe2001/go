#!/bin/bash
# this image will be 331MB
docker image rm maps || true
(cd ..; docker build -f docker/Dockerfile-with-go-runtime -t maps .)
