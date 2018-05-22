FROM golang:1.10 as build

RUN apt install make

COPY . /go/src/github.com/visheyra/prometheus-nats-gateway

RUN make -C /go/src/github.com/visheyra/prometheus-nats-gateway

ENTRYPOINT ["/go/bin/png"]
