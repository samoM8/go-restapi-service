# Go restapi service
## How to run
    $ docker run -d --name postgres-container -e TZ=UTC -p 5432:5432 -e POSTGRES_PASSWORD=postgres ubuntu/postgres:12-20.04_beta
    $ go run main.go

## Run in docker
    $ docker run -d --name postgres-container -e TZ=UTC -p 5432:5432 -e POSTGRES_PASSWORD=postgres ubuntu/postgres:12-20.04_beta
    $ docker build --tag restapi-service .
    $ docker run -p 8080:8080 -e ENV=docker --add-host=host.docker.internal:host-gateway --name restapi-service restapi-service

## Run with docker compose
    $ docker-compose up

### Swagger UI
 - http://localhost:8080/docs