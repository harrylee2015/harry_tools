#!/usr/bin/env bash

# docker 中下载 mysql
docker pull mysql:latest

#启动容器
docker run --name mysql -p 3306:3306 -v /datadir/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=fzm123!  -d mysql:latest --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

#进入容器
#docker exec -it mysql bash

#登录mysql
#mysql -u root -p
#ALTER USER 'root'@'localhost' IDENTIFIED BY 'fzm123!';

#添加远程登录用户
#CREATE USER 'harrylee'@'%' IDENTIFIED WITH mysql_native_password BY 'harrylee123!';
#GRANT ALL PRIVILEGES ON *.* TO 'harrylee'@'%';

