#!/bin/sh

docker-compose up -d

sleep 10

docker rm app-tester
docker rmi app-test
docker build -t app-test .

docker run -it --name app-tester --net host  --rm   -v "$PWD"/../../:/go/src/article-api  -w /go/src/article-api/tests/unit-tests app-test

docker-compose stop
docker-compose rm -f

