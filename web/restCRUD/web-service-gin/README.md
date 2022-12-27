DOCKER

docker build --tag docker-rest .

docker run --publish 8090:8090 docker-rest