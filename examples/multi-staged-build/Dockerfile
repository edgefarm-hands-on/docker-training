FROM alpine:3.16 as builder

RUN apk add g++ make
COPY src /app/src
COPY Makefile /app
WORKDIR /app
RUN make

FROM alpine:latest
COPY --from=builder /app/static /static
CMD [ "/static" ]