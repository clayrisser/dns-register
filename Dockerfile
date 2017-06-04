############################################################
# Dockerfile to run dns-register
# Based on Alpine
############################################################

FROM alpine:3.5

MAINTAINER Jam Risser (jamrizzi)

WORKDIR /app/

RUN apk add --no-cache tini && \
    apk add --no-cache --virtual build-deps \
    gcc \
    go \
    musl-dev \
    openssl \
    git

COPY ./ /app/.tmp/
RUN export GOPATH=/app/.tmp/ && \
    export GOBIN=/app/.tmp/bin/ && \
    cd /app/.tmp/ && \
    go get && \
    go build /app/.tmp/register.go && \
    mv /app/.tmp/register /app/register && \
    rm -rf /app/.tmp/ && \
    apk del build-deps

ENTRYPOINT ["/sbin/tini", "--", "/app/register"]
