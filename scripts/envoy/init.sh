#!/usr/bin/env sh
# Commands
docker pull envoyproxy/envoy
docker run -network host -d -v `pwd`:/cfg -w /cfg -p 10000:10000 envoyproxy/envoy -c boostrap.yaml