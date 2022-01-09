# GitHub:       https://github.com/go-dex-dev
# Twitter:      https://twitter.com/GoDexDev
# Website:      https://godex.dev/

FROM golang:1.17.6-alpine

LABEL maintainer = "https://github.com/go-dex-dev"

WORKDIR /go/src/github.com/go-dex-dev/platform

COPY . /go/src/github.com/go-dex-dev/platform
