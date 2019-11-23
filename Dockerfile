FROM golang:1.13-alpine

RUN apk update && \
    apk add \
    bash \
    git \
    build-base \
    curl \
    make \
    && rm -rf /var/cache/apk/*

WORKDIR $GOPATH/apathy

COPY . .
RUN make build

ENTRYPOINT ["/go/apathy/apathy"]