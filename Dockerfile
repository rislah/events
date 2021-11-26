FROM golang:1.17.3

ARG PROTOC_VERSION=3.19.1
ENV PROTOC_VERSION=${PROTOC_VERSION}

ARG PROTOC_GEN_GO_VERSION=1.27.1
ENV PROTOC_GEN_GO_VERSION=${PROTOC_GEN_GO_VERSION}

RUN apt-get update
RUN apt-get install -y zip

RUN wget --quiet -O /protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
RUN unzip -d /usr/local /protoc.zip

RUN go install -v google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
RUN go install -v github.com/uber/prototool/cmd/prototool@v1.10.0

ENV PROTOTOOL_PROTOC_BIN_PATH=/usr/local/bin/protoc \
    PROTOTOOL_PROTOC_WKT_PATH=/usr/local/include

RUN mkdir /.cache && chmod 777 /.cache