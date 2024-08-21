FROM ubuntu:latest AS base

# Update and install necessary packages
RUN apt-get update --fix-missing -y \
    && apt-get upgrade -y \
    && apt-get install -y ca-certificates net-tools iproute2

# Install Go 1.23
WORKDIR /usr/local/
COPY ./bin/go1.23.0.linux-amd64.tar.gz /usr/local/
RUN tar -xvzf go1.23.0.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

RUN go version

WORKDIR /opt
