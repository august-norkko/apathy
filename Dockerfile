FROM golang:1.13-alpine

RUN apk update && \
    apk add \
        bash \
        build-base \
        curl \
        make \
        git \
        && rm -rf /var/cache/apk/*

RUN mkdir /app
WORKDIR /app

ADD go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o apathy ./cmd
RUN chmod +x /app/apathy

ENTRYPOINT ["/app/apathy"]