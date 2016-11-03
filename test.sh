#! /bin/bash

docker-compose up -d

echo "Wait for RocketChat to start"
until curl -s localhost:3000 > /dev/null; do echo -n "."; sleep 1; done

go test || true

docker-compose kill
docker-compose rm -f 
