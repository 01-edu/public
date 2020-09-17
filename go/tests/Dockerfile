FROM golang:1.14.6-alpine3.12

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/01-edu/z01
COPY . src
RUN go get rc
RUN go install rc prog/...
COPY entrypoint.sh /usr/local/bin
COPY --from=mkcert-ca . /usr/local/share/ca-certificates
RUN update-ca-certificates
ENTRYPOINT ["/bin/sh", "/usr/local/bin/entrypoint.sh"]
