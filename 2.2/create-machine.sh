#!/bin/bash

for i in 1 2 3; do
    docker-machine create -d virtualbox swarm-$i
done
