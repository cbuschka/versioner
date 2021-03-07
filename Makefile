TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
VERSION ::= $(shell git describe --always --tags --dirty)
BUILD_TIME ::= $(shell date "+%Y-%m-%d_%H:%M:%S%:z")
COMMITISH ::= $(shell git describe --always --dirty)
OS ::= $(shell uname -s)
ARCH ::= $(shell uname -m)
ifeq (${GOPATH},)
	GOPATH := ${HOME}/go
endif

build:
	@cd ${TOP_DIR} && \
	mkdir -p dist/ && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a \
			-ldflags "-X github.com/cbuschka/versioner/internal/build.buildInfoVersion=${VERSION} \
					-X github.com/cbuschka/versioner/internal/build.buildInfoBuildTime=${BUILD_TIME} \
					-X github.com/cbuschka/versioner/internal/build.buildInfoCommitish=${COMMITISH} \
					-X github.com/cbuschka/versioner/internal/build.buildInfoOs=${OS} \
					-X github.com/cbuschka/versioner/internal/build.buildInfoArch=${ARCH} \
					-extldflags \"-static\"" -o dist/versioner ./cmd/versioner.go

clean:
	rm -rf ${TOP_DIR}/dist/ ${TOP_DIR}/.cache/

format:
	@cd ${TOP_DIR} && \
	gofmt -s -w .

test:
	@cd ${TOP_DIR} && \
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic ./cmd/... ./internal/...

lint:
	@cd ${TOP_DIR} && \
	go get -u golang.org/x/lint/golint && \
	${GOPATH}/bin/golint ./internal/... ./cmd/...

refresh:
	@cd ${TOP_DIR} && \
	go clean -modcache && go mod tidy 
