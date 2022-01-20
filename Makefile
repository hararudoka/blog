build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/blog/main.go

run:
	go run cmd/blog/main.go