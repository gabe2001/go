#!/bin/bash
# run go code with sample data in resources
docker run --rm --name maps -v `pwd`/../resources:/home/go/data:ro maps data/users.txt
