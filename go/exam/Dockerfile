FROM golang:1.14.6-alpine3.12

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git
RUN go get github.com/01-edu/z01
COPY --from=mkcert-ca . /usr/local/share/ca-certificates
RUN update-ca-certificates
COPY entrypoint.sh /usr/local/bin
ENTRYPOINT ["/bin/sh", "/usr/local/bin/entrypoint.sh"]
