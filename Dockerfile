FROM golang:1.13-alpine3.10
WORKDIR /go/src
ENV GO111MODULE=on
RUN apk add --no-cache \
        alpine-sdk \
        git \
    && go get github.com/pilu/fresh
EXPOSE 8080
CMD ["fresh"]