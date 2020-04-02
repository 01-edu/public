FROM debian

RUN apt-get update
RUN apt-get install -y jq curl git
COPY . /app
ENTRYPOINT ["/bin/bash", "/app/entrypoint.sh"]
