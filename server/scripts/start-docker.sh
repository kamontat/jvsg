#!/usr/bin/env bash

docker run \
  -it \
  --rm \
  -p 3333:3333 \
  jvsg/server:latest
