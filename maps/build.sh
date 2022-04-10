#!/bin/bash
docker image rm maps
docker build -f docker/Dockerfile -t maps .
