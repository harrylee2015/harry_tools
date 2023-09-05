#!/bin/bash

docker run -d -p 2181:2181 ubuntu/zookeeper:edge

sleep 5

docker run -d --name kafka-container -e TZ=UTC -p 9092:9092 -e ZOOKEEPER_HOST=host.docker.internal ubuntu/kafka:3.1-22.04_beta

