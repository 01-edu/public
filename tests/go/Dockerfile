FROM golang:1

RUN go get golang.org/x/tools/cmd/goimports
RUN go get github.com/01-edu/z01
RUN go get github.com/01-edu/public/rc
COPY . /app
ENTRYPOINT ["/bin/bash", "/app/entrypoint.sh"]
