run:
	go run ./cmd/main.go
build:
	go build -o ./bin/server ./cmd/main.go
start:
	./bin/server -b 0.0.0.0
tidy:
	go mod tidy
