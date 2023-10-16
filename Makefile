# --name is the name of container
# can use docker logs to see the postgres conatiner logs
# the db name will same as username if not mentioned explicitly, here it will be root.
postgres:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:alpine

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root chat-db

dropdb:
	docker exec -it postgresdb dropdb --username=root --owner=root chat-db

migrate-up:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/chat-db?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/chat-db?sslmode=disable" -verbose up

sqlc:
	sqlc generate

# -v verbose logs, -cover for code coverage data, ./... to run unit test in all the packages
test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: createdb dropdb migrate-down migrate-up sqlc server