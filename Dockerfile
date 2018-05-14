FROM golang:1.10 as build

RUN apt install make

COPY . /go/src/github.com/visheyra/prometheus-nats-gateway

RUN make -C /go/src/github.com/visheyra/prometheus-nats-gateway

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=build /go/bin/png png

ENTRYPOINT ["/app/png"]
