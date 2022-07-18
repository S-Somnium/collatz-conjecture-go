FROM golang:latest
RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN cd /build && git clone 