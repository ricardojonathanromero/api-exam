# 130 MB
FROM golang:1.14.15-alpine AS build
WORKDIR /app/
RUN apk add --no-cache git ca-certificates
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o api-exam .

# 30MB
FROM alpine as final
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=build /app/api-exam ./
RUN chown nobody:nobody ./api-exam
USER nobody
ENTRYPOINT ["./api-exam"]