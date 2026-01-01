test:
	go test ./...

run:
	go run main.go -day 1 -part 1 -all false
	go run main.go -day 1 -part 2 -all false

build:
	go build -o advent2025 main.go