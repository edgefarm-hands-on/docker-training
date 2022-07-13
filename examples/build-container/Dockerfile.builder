FROM alpine:3.16 as builder

RUN apk add g++ make
WORKDIR /app