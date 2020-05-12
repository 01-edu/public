FROM golang:1-alpine

RUN apk add git
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/01-edu/z01
RUN go get github.com/01-edu/public/rc
COPY . /app
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
