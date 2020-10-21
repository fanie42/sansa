FROM golang:latest AS build-env

ADD . /go/src/sansa
RUN cd /go/src/sansa \
    && ls \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags "-w -X main.docker=true" ./cmd/daq/main.go

FROM scratch
LABEL maintainer="Stephanus Schoeman <sschoeman94@gmail.com>"

ENV LOCAL_BASE_PREFIX="marl111_"
ENV LOCAL_BASE_FORMAT="20060102_15"
ENV LOCAL_BASE_POSTFIX="0000"
ENV LOCAL_BASE_EXTENSION="dat"

ENV LOCAL_PATH_LOCATION="./data"
ENV LOCAL_PATH_FORMAT="./2006/01/02"

ENV SERIAL_NAME="COM5"
ENV SERIAL_BAUD="19200"

ENV NATS_URL="nats://172.18.30.100:4222"
ENV NATS_CLUSTER_ID="marion"
ENV NATS_CLIENT_ID="marl111_daq"

COPY --from=build-env /go/src/sansa/main /
CMD ["/main"]