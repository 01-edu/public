FROM node:14-alpine

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git
COPY --from=mkcert-ca . /usr/local/share/ca-certificates
RUN update-ca-certificates
WORKDIR /app
COPY . .
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
