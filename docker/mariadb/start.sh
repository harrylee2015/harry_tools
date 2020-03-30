#!/usr/bin/env bash

docker run --name my-mariadb -v /datadir/mariadb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=mysecret -d mariadb:latest --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
