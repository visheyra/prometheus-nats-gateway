#
# GLOBAL VARIABLES
#

VERSION := "0.1.0"
TARGET := $(GOPATH)/bin/png
IMAGE_NAME := "visheyra/prometheus-nats-gateway"
IMAGE_TAG := "latest"
DOCKER_ARGS := ""
PACKAGE := github.com/visheyra/prometheus-nats-gateway

GO := $(shell which go)
ifeq ($(GO),)
	$(error Could not find go)
endif

export VERBOSE ?= 1
Q = $(if $(filter 1,$VERBOSE),,@)
H = $(shell printf "\033[34;1m=>\033[0m")

#
# Default target
#

all: $(TARGET)

$(TARGET): depend build

#
# Dependancy targets
#

depend: depend.tools depend.vendor

depend.tools:
	$(info $(H) get dependancy tools)
	$(Q) curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

depend.vendor:
	$(info $(H) get dependancy packages)
	$(Q) dep ensure -v

.PHONY: depend deptoolend.tools depend.vendor

#
# Build target
#

build: build.app

build.app:
	$(info $(H) build app)
	$(Q) $(GO) build -v -o $(TARGET)

build.docker:
	$(info $(H) build docker image)
	$(Q) docker build $(ARGS) -t $(IMAGE_NAME):$(IMAGE_TAG) .

.PHONY:	build build.app build.docker

#
# Clean target
#

clean: clean.app clean.vendor

clean.app:
	$(info $(H) deleting binary)
	$(Q) rm -rf $(TARGET)

clean.vendor:
	$(info $(H) deleting vendor directory)
	$(Q) rm -rf vendor
