FROM golang:1.13-alpine

RUN apk update && \
    apk add \
    bash \
    build-base \
    curl \
    make \
    tzdata \
    git \
    && rm -rf /var/cache/apk/*

RUN adduser -D -g '' u
WORKDIR $GOPATH/src/apathy
USER u

# Install modules
ADD go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .

# Build binary
RUN GOOS=linux go build -a -o $GOPATH/bin/apathy .
RUN chmod +x /go/bin/apathy

# Run binary
CMD ["/go/bin/apathy"]