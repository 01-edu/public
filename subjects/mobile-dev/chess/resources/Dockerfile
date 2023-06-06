FROM golang:1.19.1-buster AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /app/main .


FROM ubuntu:20.04

WORKDIR /app
COPY --from=builder /app/main .
CMD ["/app/main"]
