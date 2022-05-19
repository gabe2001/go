FROM hub.docker.hpecorp.net/hub/golang:1.18-alpine
RUN addgroup -S go && adduser -S go -G go
USER go:go
WORKDIR /home/go
CMD mkdir data
COPY ./src .
COPY go.mod .
RUN go build -o ./bin/maps
ENTRYPOINT ["./bin/maps"]
