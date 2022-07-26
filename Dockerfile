FROM golang:1.18.4 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:latest AS deploy
COPY --from=builder /app .
CMD ["./myapp"]