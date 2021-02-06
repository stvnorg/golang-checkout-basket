FROM golang:1.14.15-alpine3.11

WORKDIR /lana/

COPY go.mod go.sum api.go basket.go discounts.go ./

RUN GOOS=linux go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
ENV DB_NAME=basket_db

COPY --from=0 /lana/app .

CMD ["./app"]
