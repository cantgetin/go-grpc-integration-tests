# Go gRPC integration tests

Basic gRPC service written in Go with following proto APIs:
1) Users - GetUser, ListUsers
2) Books - GetBook, ListBooks


Service has simple integration tests of these endpoints using dockertest and testify suite \
Migrations with scheme and sample data being applied then tests start \
Tests compare data from migrations with actual response from service 

## Run tests
```shell
make integration-tests
```

## Run app
```shell
make up
make migrate
make run #or make run-docker
```