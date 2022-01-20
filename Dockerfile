FROM golang:alpine
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 go build -o /usr/bin/main cmd/blog/main.go
ENTRYPOINT ["main"]