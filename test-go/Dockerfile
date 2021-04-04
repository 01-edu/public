FROM docker.01-edu.org/golang:1.16.3-alpine3.13

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git

RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/01-edu/rc

WORKDIR /piscine-go
RUN go mod init piscine-go
RUN go get github.com/01-edu/z01@v0.1.0

WORKDIR /public/test-go
COPY go.* ./
RUN go mod download
COPY lib lib
COPY solutions solutions
COPY tests tests
RUN go install $(grep -rl ChallengeMain ./tests | rev | cut -d/ -f2- | rev)

RUN rm -rf /piscine-go

COPY entrypoint.sh /usr/local/bin
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
