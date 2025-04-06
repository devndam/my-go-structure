run:
	go run main.go

migrate:
	go run tools/migrate.go

test:
	go test ./...