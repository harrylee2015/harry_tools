#!/usr/bin/env bash

#拉取镜像
docker pull elcolio/etcd:latest

#启动服务

docker run \
  -d \
  -p 2379:2379 \
  -p 2380:2380 \
  -p 4001:4001 \
  -p 7001:7001 \
  -v /data/etcd:/data \
  --name some-etcd \
  elcolio/etcd:latest \
  -name some-etcd \
