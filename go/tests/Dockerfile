FROM docker.01-edu.org/golang:1.16.3-alpine3.13

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git

RUN go get golang.org/x/tools/cmd/goimports

WORKDIR /piscine-go
RUN go mod init piscine-go
RUN go get github.com/01-edu/z01@v0.1.0

WORKDIR /public/go/tests
COPY go.* ./
RUN go mod download
COPY lib lib
COPY prog prog
RUN go get github.com/01-edu/rc
RUN go install $(grep -rl ChallengeMain ./prog | rev | cut -d/ -f2- | rev)

RUN rm -rf /piscine-go

COPY entrypoint.sh /usr/local/bin
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
