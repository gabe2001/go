FROM hub.docker.hpecorp.net/hub/alpine:3.13.6
RUN addgroup -S go && adduser -S go -G go
USER go:go
WORKDIR /home/go
CMD mkdir data
COPY maps .
ENTRYPOINT ["./maps"]
