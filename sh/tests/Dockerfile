FROM alpine

RUN apk add jq curl git
COPY . /app
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
