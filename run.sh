#!/bin/bash
docker kill $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker build -t groupcache_example .
docker run --rm -t -p 8080:8080 -e ETCD_ADDR="http://${etcd}:4001" -e PUBLIC_PORT="8080" groupcache_example &
docker run --rm -t -p 8081:8080 -e ETCD_ADDR="http://${etcd}:4001" -e PUBLIC_PORT="8081" groupcache_example
