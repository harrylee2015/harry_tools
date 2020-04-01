#!/usr/bin/env bash
# ----------------------------------------------------------------------
# name:         docker-install.sh
# version:      1.0
# createTime:   2020-02-03
# description:  安装docker及docker-compose，运行前提确保机器能够连上外网
# author:       harry lee
# ----------------------------------------------------------------------
export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
# Check the network status
NET_NUM=`ping -c 4 www.baidu.com |awk '/packet loss/{print $6}' |sed -e 's/%//'`
if [ -z "$NET_NUM" ] || [ $NET_NUM -ne 0 ];then
        echo "Please check your internet"
        exit 1
fi
sudo docker version
if [ $? -eq 0 ];then
   echo "This machine has installed docker!"
else
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
   sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
   sudo apt-get update
   sudo apt-get -y -o Dpkg::Options::="--force-confnew" install docker-ce  --allow-unauthenticated
fi
sudo docker-compose version
if [ $? -eq 0 ];then
   echo "This machine has installed docker-compose!"
else
   sudo curl -L https://github.com/docker/compose/releases/download/1.25.4/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
   sudo chmod +x /usr/local/bin/docker-compose
fi
