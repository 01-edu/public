FROM buildkite/puppeteer:7.1.0

WORKDIR /app
COPY dom .
COPY subjects ./subjects
ENTRYPOINT ["/app/entrypoint.sh"]
