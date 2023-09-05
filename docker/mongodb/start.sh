#!/bin/bash
docker run -p 27017:27017 --name some-mongo -v $(pwd)/datadir:/data/db -d  --privileged=true mongo --wiredTigerCacheSizeGB 0.4 --auth

