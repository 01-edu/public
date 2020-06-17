FROM buildkite/puppeteer:latest

ENV GIT_TERMINAL_PROMPT=0
# RUN apk add --no-cache git
WORKDIR /app
COPY ./puppeteer .
COPY ./subjects ./subjects
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
