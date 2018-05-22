# prometheus-nats-gateway

Publish prometheus metrics into NATS

[![Build Status](https://travis-ci.org/visheyra/prometheus-nats-gateway.svg?branch=master)](https://travis-ci.org/visheyra/prometheus-nats-gateway)

## Description

This tool is only compliant with prometheus 2.0 instances as it uses [remote write](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#%3Cremote_write%3E) feature to collect timeseries from prometheus.
The nats client used in this project is the [vanilla one](https://github.com/nats-io/go-nats)

## Distribution

### Docker

You just have to pull **visheyra/prometheus-nats-gateway**

### Build it yourself

1. Clone this repo in the `${GOPATH}/src/github.com/visheyra/prometheus-nats-gateway` directory
2. Download the [dep tool](https://github.com/golang/dep)
3. go in the `src` folder of this repo
4. run `go dep ensure`
5. run go install github.com/visheyra/prometheus-nats-exporter
6. the binary will be in `${GOPATH}/bin`

## Running

```
tool that listen prometheus 2.0 events translate them to json, then publish them to nats

Usage:
  png [command]

Available Commands:
  help        Help about any command
  start       Start the tool

Flags:
  -f, --forward string   address of the remote nats endpoint (default "http://localhost:4222")
  -h, --help             help for png
  -l, --listen string    listen address of the prometheus receiver endpoint (default ":8080")

Use "png [command] --help" for more information about a command.
```

## Authentication

| Type | Support | Description |
|:---:|:---:|:---:|
|certificates| no | not implemented |
| user creds | partial | only when supplied in the URI (eg: nats://user:pass@server:port)|
| plain (no auth) | yes | - |

## TODO

* [ ] Support user sefined NATS subject
* [ ] Support authentication methods (current implementation support only plain connections)
* [ ] Add a /metrics endpoint for telemetry
