#!/usr/bin/env bash
make build
tar -czvf heartbeat.tar.gz dist/heartbeat
rsync -azv heartbeat.tar.gz root@centaurwarchief.com:/home/ubuntu/heartbeat.centaurwarchief.com
rm -rf dist/heartbeat
rm -rf heartbeat.tar.gz
ssh root@centaurwarchief.com <<'C'
  cd /home/ubuntu/heartbeat.centaurwarchief.com
  tar -zxvf heartbeat.tar.gz
  mv dist/heartbeat heartbeat
  rm -rf dist
  rm -f heartbeat.tar.gz
  docker stop heartbeat.centaurwarchief.com |xargs docker rm >/dev/null 2>&1
  docker run \
    -d \
    --name heartbeat.centaurwarchief.com \
    --net=host \
    -v `pwd`:/heartbeat \
    -w /heartbeat \
    busybox:latest ./heartbeat
C
