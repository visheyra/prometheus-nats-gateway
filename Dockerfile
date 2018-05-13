FROM golang:1.10 as build

COPY src /go/src/github.com/visheyra/prometheus-nats-gateway

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR /go/src/github.com/visheyra/prometheus-nats-gateway

RUN ${GOPATH}/bin/dep ensure -v

RUN go install -v github.com/visheyra/prometheus-nats-gateway

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=build /go/bin/prometheus-nats-gateway png

ENTRYPOINT ["/app/png"]
