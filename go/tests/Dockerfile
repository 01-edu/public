FROM golang:1-alpine

RUN apk add --no-cache git
RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/01-edu/z01
WORKDIR src
COPY . .
COPY entrypoint /usr/local/bin
RUN go install -trimpath -ldflags="-s -w" rc
RUN find -mindepth 2 -maxdepth 2 -name '*_correct' -type d|xargs go install -trimpath -ldflags="-s -w"
RUN find -mindepth 2 -maxdepth 2 -name '*_correct' -type d|cut -d/ -f1-2|xargs go install -trimpath -ldflags="-s -w"
RUN find -mindepth 2 -maxdepth 2 -name '*_correct' -type d|cut -d/ -f1-2|xargs rm -rf
# compare performance with and without :
#   ENV CGO_ENABLED 0
#   -trimpath -ldflags="-s -w"
#   -i
#   -a (to rebuild everything and prepare .a files)
# git clone in Dockerfile to make it work anywhere (no COPY)
