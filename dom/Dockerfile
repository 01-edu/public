FROM buildkite/puppeteer:5.2.1

WORKDIR /app
COPY dom .
COPY subjects ./subjects
ENTRYPOINT ["/app/entrypoint.sh"]
