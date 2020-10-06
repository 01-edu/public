FROM debian:10.5-slim

RUN apt-get update
RUN apt-get -y install jq curl git
COPY . /app
ENTRYPOINT ["/app/entrypoint.sh"]
