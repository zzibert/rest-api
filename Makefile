build:
	go build -o bin/main main.go

run:
	go run main.go

start: up run

up:
	cd docker/ && docker-compose up -d

down:
	cd docker/ && docker-compose down

test:
	go test -check.vv

