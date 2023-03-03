CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
LINTVER=v1.51.2
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=gocloudcamp

testcli: build_testcli
	@bin/testcli

server: build_server
	@bin/server

client: build_client
	@bin/client

up:
	@echo 'waiting for containers to start'
	@cd deploy/postgres && docker compose up -d
	@sleep 10
	@echo 'create database playlist'
	@cd deploy/postgres && docker-compose exec postgres psql -U user -c 'CREATE DATABASE playlist'

down:
	@echo 'waiting for containers to stop'
	@cd deploy/postgres && docker compose down

test:
	go test ./...

precommit: build_testcli build_server build_client test lint
	@echo "OK"

gen: clean
	@protoc --proto_path=api --go_out=pkg --go-grpc_out=pkg api/*.proto

clean:
	-@rm pkg/plservice/*.go
	-@rm -rf bin

build_testcli: bindir
	@go build -o ${BINDIR}/testcli ${PACKAGE}/cmd/testcli

build_server: bindir gen
	@go build -o ${BINDIR}/server ${PACKAGE}/cmd/server

build_client: bindir gen
	@go build -o ${BINDIR}/client ${PACKAGE}/cmd/client

lint: install-lint
	@${LINTBIN} run

bindir:
	@mkdir -p ${BINDIR}

install-lint: bindir
	@test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})
