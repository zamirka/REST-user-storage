FROM golang:1.11-alpine3.8

LABEL author="Zamir Musaev"

RUN apk add --update alpine-sdk
RUN apk add git

EXPOSE 8080

