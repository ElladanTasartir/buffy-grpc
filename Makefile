PROTO_OUT ?= gen/go
OUT_DIR ?= bin/
DIR ?= $(shell pwd)

PROTOC_CMD ?= docker run --rm -v $(DIR):$(DIR) -w $(DIR) thethingsindustries/protoc

.PHONY: compile-proto-go
compile-proto-go: ## Compile the protobuf contracts to generated Go code
	@sudo rm -rf $(PROTO_OUT)
	@mkdir -p $(PROTO_OUT)
	@find proto -name "*.proto" -type f -exec $(PROTOC_CMD) -Iproto/ \
         --go_opt=paths=source_relative \
	     --go_out=plugins=grpc:$(PROTO_OUT) {} \;

.PHONY: build
build:
	@go build -o $(OUT_DIR)/main cmd/buffy-grpc/main.go

.PHONY: run
run: build
	@./$(OUT_DIR)/main
