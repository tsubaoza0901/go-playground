FROM  golang:1.18.4

WORKDIR /go/src

# RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN go install github.com/99designs/gqlgen@latest