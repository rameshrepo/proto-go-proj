FROM ubuntu:latest

ENV GO_VERSION=1.20.6

RUN apt-get update && apt-get install -y wget git gcc curl protobuf-compiler make
RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz"

RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-amd64.tar.gz"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

WORKDIR $GOPATH