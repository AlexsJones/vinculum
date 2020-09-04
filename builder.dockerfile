FROM  onosproject/protoc-go as base
ADD . /src/
WORKDIR /src
RUN protoc --proto_path=pkg/proto pkg/proto/*.proto --go_out=plugins=grpc:.

FROM golang:1.13.1-alpine3.10 as builder
# System setup
RUN apk update && apk add git curl build-base autoconf automake libtool
COPY --from=base /src/ /src/
WORKDIR /src
RUN go mod download && \
    go get github.com/mitchellh/gox
RUN make
