FROM node:14-alpine

ENV GIT_TERMINAL_PROMPT=0
RUN apk add --no-cache git
WORKDIR /app
COPY . .
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
