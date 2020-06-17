FROM buildkite/puppeteer:latest

ENV GIT_TERMINAL_PROMPT=0
RUN apt-get update
RUN apt-get install -y git
WORKDIR /app
COPY ./dom .
COPY ./subjects ./subjects
RUN ls -la
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
