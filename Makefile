proto:
	protoc --go_out=. --go-grpc_out=. api/books/books.proto
	protoc --go_out=. --go-grpc_out=. api/users/users.proto

migrate:
	sql-migrate up -env="local"

integration-tests:
	go test ./tests/... -cover -v -count=1

up:
	docker-compose up

run:
	go run cmd/main.go || exit 1

run-docker:
	docker build .
