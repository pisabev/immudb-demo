#!/bin/sh

HOST=immudb:immudb@localhost:8080

# Log Insert
curl -X POST http://$HOST/api/v1/log \
  -H 'Content-type: application/json' \
  -d '{"data":"random data","source":"random source"}'

# Log Insert batch
curl -X POST http://$HOST/api/v1/logs \
  -H 'Content-type: application/json' \
  -d '[{"data":"random data 1","source":"random source 1"},{"data":"random data 2","source":"random source 2"}]'

# Logs fetch all logs
curl -X GET http://$HOST/api/v1/log

# Logs fetch last 2 logs
curl -X GET http://$HOST/api/v1/log?lastx=2