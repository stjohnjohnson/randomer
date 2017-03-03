FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/randomer /opt/bin/randomer

CMD ["/opt/bin/randomer"]
