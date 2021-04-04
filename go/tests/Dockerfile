FROM docker.01-edu.org/golang:1.16.3-alpine3.13

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git

RUN go get golang.org/x/tools/cmd/goimports

RUN mkdir -p /piscine-go
RUN touch /piscine-go/go.mod

WORKDIR /public/go/tests
COPY go.* ./
RUN go mod download
COPY func func
COPY lib lib
COPY prog prog
RUN go get github.com/01-edu/rc
RUN go install ./prog/...
RUN rm -rf /piscine-go

COPY entrypoint.sh /usr/local/bin
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
