FROM golang:alpine as builder
RUN apk --update add --no-cache bash curl coreutils
WORKDIR /app
ADD . .
RUN go build -o cronjob main.go

FROM alpine as prod
WORKDIR /app
COPY --from=builder /app/cronjob /app/cronjob

