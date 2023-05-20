proto:
	protoc --go_out=. --go-grpc_out=. api/books/books.proto
	protoc --go_out=. --go-grpc_out=. api/users/users.proto

migrate:
	sql-migrate up -env="local"