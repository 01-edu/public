FROM golang:1-alpine

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git nano
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/01-edu/z01
WORKDIR src
COPY entrypoint.sh /usr/local/bin
COPY tests .
RUN go install rc
RUN find -mindepth 2 -maxdepth 2 -name '*_correct' -type d|xargs go install
RUN find -mindepth 2 -maxdepth 2 -name '*_correct' -type d|cut -d/ -f1-2|xargs go install
RUN find -mindepth 2 -maxdepth 2 -name '*_correct' -type d|cut -d/ -f1-2|xargs rm -rf
WORKDIR /tmp
USER 1000
ENTRYPOINT ["/bin/sh", "/usr/local/bin/entrypoint.sh"]
# compare performance with and without :
#   ENV CGO_ENABLED 0
#   -trimpath -ldflags="-s -w"
#   -i
#   -a (to rebuild everything and prepare .a files)
# git clone in Dockerfile to make it work anywhere (no COPY)
