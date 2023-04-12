GO=go
GOMODULE=meta-egg-layout
debug?=0 # 0: release, 1: support dlv debug

LDFLAGS += -X "${GOMODULE}/pkg/version.BuildTS=$(shell date -u '+%Y-%m-%d %I:%M:%S')"
LDFLAGS += -X "${GOMODULE}/pkg/version.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "${GOMODULE}/pkg/version.GitBranch=$(shell git rev-parse --abbrev-ref HEAD)"
LDFLAGS += -X "${GOMODULE}/pkg/version.GitTag=$(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)"
LDFLAGS += -X "${GOMODULE}/pkg/version.GitDirty=$(shell test -n "` + "`" + `git status --porcelain` + "`" + `" && echo "true")"
LDFLAGS += -X "${GOMODULE}/pkg/version.Debug=$(shell if [ ${debug} -eq 1 ]; then echo "true"; fi)"

ifeq ($(debug), 1)
	GC_FLAGS:=-gcflags "all=-N -l"
else
	GO_LDFLAGS += -s -w
	GO_TRIMPATH = -trimpath
endif


TARGETS := $(shell ls cmd)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	go get github.com/google/wire/cmd/wire@latest
	go get github.com/golang/mock/mockgen@v1.6.0
	go mod tidy

.PHONY: pb
# generate api proto
pb:
	protoc \
	  --proto_path ./proto \
	  --proto_path ./third_party/proto \
	  --go_out=./api \
	  --go-grpc_out=./api \
	  --validate_out=lang=go:./api \
	  ./proto/*.proto

generate: pb
# generate
generate:
	go get github.com/google/wire/cmd/wire@latest
	go get github.com/golang/mock/mockgen@v1.6.0
	go generate ./...
	go mod tidy

.PHONY: vet
# vet code
vet:
	go vet ./...

build: vet
# build binary
build: $(TARGETS)

$(TARGETS):
	${GO} build -ldflags '${GO_LDFLAGS} ${LDFLAGS}' ${GC_FLAGS} ${GO_TRIMPATH} -o build/bin/ ${GOMODULE}/cmd/$@

.PHONY: clear
# clear binary
clear:
	rm -f build/bin/*

.PHONY: run
# run binary
run:
	./build/bin/%%PROJECT-NAME%% -config ./configs/conf-local.yml

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help