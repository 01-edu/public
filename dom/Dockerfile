FROM buildkite/puppeteer:5.2.1

ENV GIT_TERMINAL_PROMPT=0
RUN apt-get update
RUN apt-get install -y git
COPY --from=mkcert-ca . /usr/local/share/ca-certificates
RUN update-ca-certificates
WORKDIR /app
COPY dom .
COPY subjects ./subjects
ENTRYPOINT ["/bin/sh", "/app/entrypoint.sh"]
