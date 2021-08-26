# 130 MB
FROM golang:1.14.15-alpine as build


# 30MB
FROM alpine as final
WORKDIR /app/
USER noroot
ENTRYPOINT ["./demo"]