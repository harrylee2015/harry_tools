#!/bin/bash
docker run -d --hostname my-rabbit --name some-rabbit -p 5671:5617 -p 5672:5672 -p 4369:4369 -p 15671:15671 -p 15672:15672 -p 25672:25672  rabbitmq:3-management
