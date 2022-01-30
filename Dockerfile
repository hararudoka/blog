FROM golang:alpine as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

FROM alpine

WORKDIR /app

COPY --from=builder /build/blog .
COPY --from=builder /build/view view

EXPOSE 80/tcp 443/tcp

ENTRYPOINT ["/app/blog"]