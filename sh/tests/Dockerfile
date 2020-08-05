FROM debian:10.5-slim

RUN apt-get update
RUN apt-get -y install jq curl git
COPY --from=mkcert-ca . /usr/local/share/ca-certificates
RUN update-ca-certificates
COPY . /app
ENTRYPOINT ["/bin/bash", "/app/entrypoint.sh"]
