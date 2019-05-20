#!/bin/bash

docker-machine create -d virtualbox swarm-1

eval "$(docker-machine env swarm-1)"

docker swarm init --advertise-addr $(docker-machine ip swarm-1)

docker service create \
  --name=visualizer \
  --publish=8000:8080/tcp \
  --constraint=node.role==manager \
  --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock \
  dockersamples/visualizer
