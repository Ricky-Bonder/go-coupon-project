# build stage
FROM golang:latest AS builder

RUN apt update
RUN apt-get install -y ca-certificates && update-ca-certificates
RUN apt-get install -y \
    git \
    gcc \
    libc-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/coupon-service

COPY . .

WORKDIR /go/src/coupon-service/cmd/coupon_service
RUN go build -o main .

ENTRYPOINT ./main
EXPOSE 8080
