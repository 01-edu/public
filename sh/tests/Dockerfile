FROM debian:stable-slim

RUN apt-get update
RUN apt-get -y install jq curl tree apt-utils
RUN apt-get -y install bc

WORKDIR /app/assets/superhero
RUN curl --remote-name --location https://demo.01-edu.org/assets/superhero/all.json

WORKDIR /tmp/installation

RUN curl --remote-name --location https://github.com/caddyserver/caddy/releases/download/v2.3.0/caddy_2.3.0_linux_amd64.tar.gz
RUN tar xf caddy_2.3.0_linux_amd64.tar.gz
RUN mv caddy /usr/local/bin
RUN apt-get -y install libcap2-bin
RUN setcap 'cap_net_bind_service=+ep' /usr/local/bin/caddy

WORKDIR /app

RUN rm -rf /tmp/installation

COPY . /app
ENTRYPOINT ["/app/entrypoint.sh"]
