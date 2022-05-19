#!/bin/bash
# this image will *ONLY* be 7.49MB (vs 331MB with go SDK)
docker image rm maps || true
(cd ..; go build; docker build -f docker/Dockerfile -t maps .)
