FROM golang:1.24.1-alpine as builder

RUN mkdir app

ENV DB_URL=""

WORKDIR /app
COPY ./ .
run go build -o main .

EXPOSE 8080

FROM alpine:3
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]