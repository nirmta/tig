# Dev
FROM golang:1.26.5-alpine AS dev

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git openssh-client

RUN go install github.com/goreleaser/goreleaser/v2@v2.17.1

CMD ["sleep", "infinity"]

# Test
FROM golang:1.26.5-alpine AS test

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git openssh-client

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "test", "-v", "./..."]

# Builder
FROM golang:1.26.5-alpine AS build

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git openssh-client

COPY . .

RUN go install github.com/goreleaser/goreleaser/v2@v2.17.1

CMD ["goreleaser", "release", "--snapshot"]

# Releaser
FROM golang:1.26.5-alpine AS release

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git openssh-client

RUN go install github.com/goreleaser/goreleaser/v2@v2.17.1

CMD ["goreleaser", "release", "--clean"]
