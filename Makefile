build:
	go build -o bin/uncomment cmd/uncomment/main.go

run:
	go run cmd/main.go

test:
	go test

install:
	go install ./cmd/uncomment
