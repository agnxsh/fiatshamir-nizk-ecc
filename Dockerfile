FROM golang:1.18-alpine as builder

WORKDIR /fiatshamir-nizk-ecc

COPY . .

RUN go build


