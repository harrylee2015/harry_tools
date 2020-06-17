#!/usr/bin/env bash

docker run --name nigix -v $(pwd)/nginx.conf:/etc/nginx/nginx.conf:ro -d nginx

