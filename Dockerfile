FROM ubuntu:latest AS base

# Update and install necessary packages
RUN apt-get update --fix-missing -y \
    && apt-get upgrade -y \
    && apt-get install -y net-tools iproute2 wget

# Install Go 1.23
WORKDIR /usr/local/
COPY ./bin/go1.23.0.linux-amd64.tar.gz /usr/local/
RUN tar -xvzf go1.23.0.linux-amd64.tar.gz
ENV GOPATH=/opt/
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH

RUN go version

WORKDIR /opt
