#!/usr/bin/env bash

#拉取镜像
docker pull redis:latest

#运行容器
docker run -itd --name redis  -v /datadir/redis:/data  -p 6379:6379 redis

#docker exec -it redis /bin/bash




