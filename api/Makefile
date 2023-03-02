OUT        = ../../pkg/grpc_client
SWAGGER    = ../api
INCLUDEDIR = $(CURDIR)/includes
PROTOS     = $(shell find . -type f -name '*.proto')

BUF_VERSION:=v1.9.0
SWAGGER_UI_VERSION:=v4.15.5


all: gen

.PHONY: gen
gen: gen/proto

gen/proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate
gen/swagger-ui:
	SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) ./scripts/generate-swagger-ui.sh