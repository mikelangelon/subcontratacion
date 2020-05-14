FROM golang:alpine as builder
RUN apk --update add --no-cache bash curl coreutils
WORKDIR /app
ADD . .
RUN go build -o coolapp cmd/supercoolservice/main.go

FROM alpine as prod
WORKDIR /app
COPY --from=builder /app/coolapp /app/coolapp
EXPOSE 8080
CMD ["./coolapp"]
