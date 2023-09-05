#!/bin/bash 


docker run  -p 8086:8086 --name influx18     -v influxdb:/var/lib/influxdb   -d    influxdb:1.8
