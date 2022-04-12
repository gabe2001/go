#!/bin/bash
docker image rm maps || true
docker build -f docker/Dockerfile -t maps .
