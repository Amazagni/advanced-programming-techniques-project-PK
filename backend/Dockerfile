FROM golang:1.23 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM debian:bookworm-slim

EXPOSE 50001

COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
