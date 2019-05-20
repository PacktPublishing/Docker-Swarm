#!/bin/bash

docker-machine create -d virtualbox swarm-1

eval "$(docker-machine env swarm-1)"

docker swarm init --advertise-addr $(docker-machine ip swarm-1)


docker pull jenkins:alpine
docker pull nginx:alpine
docker pull golang:alpine
docker pull mongo:3.2.15
docker pull albertogviana/names-demo:3.0.0
docker pull dockersamples/visualizer


docker volume create -d local my-volume

docker network create -d overlay --opt encrypted names-demo
