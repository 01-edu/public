FROM debian:stable-slim

RUN apt-get update
RUN apt-get -y install jq curl git
COPY . /app
ENTRYPOINT ["/bin/bash", "/app/entrypoint.sh"]
