# Dev
FROM golang:1.26.5-alpine AS dev

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git

RUN go install
RUN go install github.com/goreleaser/goreleaser/v2@v2.17.1

CMD ["sleep", "infinity"]

# Builder
FROM golang:1.26.5-alpine AS build

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git

COPY . .

RUN go install
RUN go install github.com/goreleaser/goreleaser/v2@v2.17.1

# Fix --clean problem
CMD ["goreleaser", "release", "--snapshot"]

# Releaser
FROM golang:1.26.5-alpine AS release

WORKDIR /app

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git

RUN go install
RUN go install github.com/goreleaser/goreleaser/v2@v2.17.1

CMD ["goreleaser", "release", "--clean"]
