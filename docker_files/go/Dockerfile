FROM  golang:1.17-alpine

WORKDIR /go/app

ENV GO111MODULE=on

RUN apk update && apk add graphviz
# RUN apk add sensible-utils

# RUN go get bitbucket.org/liamstask/goose/cmd/goose