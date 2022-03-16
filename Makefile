# ########################################################## #
# Makefile for Golang Project
# Includes cross-compiling, installation, cleanup
# ########################################################## #

# Default Goal of the makefile is to show the help
.DEFAULT_GOAL := help

# Sets default shell to Bash
#SHELL := /bin/bash

# Check for required command tools to build or stop immediately
#EXECUTABLES = go find pwd awk docker
#K := $(foreach exec,$(EXECUTABLES),\
#        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

API_DIR:=$(ROOT_DIR)/api
PROTO_DIR:=$(API_DIR)/proto
SRC_DIR:=$(ROOT_DIR)
OUT_DIR:=$(ROOT_DIR)/bin

# Indicates that the following targets have no physical files
.PHONY: all proto clean build help

proto: ## Generate protobuf Go stubs
	cd $(PROTO_DIR) && protoc --go_out=plugins=grpc:$(API_DIR)/ hermit-shell-grpc/grpc/*.proto

all: proto build ## Runs everything

clean: ## Clean generated sources
	rm -rf $(OUT_DIR)

build: ## Build the application
	cd $(SRC_DIR) && go mod tidy -v && go mod vendor
	cd $(SRC_DIR) && go build -mod=vendor -o $(OUT_DIR)/main

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

