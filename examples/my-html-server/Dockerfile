FROM golang:1.18.3-alpine3.16

COPY main.go /src/main.go
COPY go.mod /src/go.mod
WORKDIR /src
RUN go build -o /app

ENTRYPOINT [ "/app"]