ARG BASE_IMAGE=registry.opensuse.org/cloud/platform/quarks/sle_15_sp1/quarks-operator-base:latest

################################################################################
FROM golang:1.15.1 AS build
ARG GOPROXY
ENV GOPROXY $GOPROXY

WORKDIR /go/src/code.cloudfoundry.org/quarks-gora

COPY . .
RUN go build && \
    cp -p quarks-gora /usr/local/bin/quarks-gora

################################################################################
FROM $BASE_IMAGE
LABEL org.opencontainers.image.source https://github.com/cloudfoundry-incubator/quarks-gora
COPY --from=build /usr/local/bin/quarks-gora /usr/local/bin/quarks-gora

ENTRYPOINT ["/usr/local/bin/quarks-gora"]
