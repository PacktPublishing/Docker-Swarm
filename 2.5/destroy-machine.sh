#!/bin/bash

for i in 1 2 3; do
    docker-machine rm -f swarm-$i
done
