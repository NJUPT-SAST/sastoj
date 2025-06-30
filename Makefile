GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
PROJECTS=admin user-contest public-auth gojudge-server

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\git.exe,bin\bash.exe,$(shell where git))
    Git_Bash="$(subst \,/,$(subst cmd\git.exe,bin\bash.exe,$(shell where git)))"
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
	APP_PROTO_FILES=$(shell $(Git_Bash) -c "find app -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
	APP_PROTO_FILES=$(shell find app -name *.proto)
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./app \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./app \
	       $(APP_PROTO_FILES)

.PHONY: api
# generate protobuf api go code
api:
	@cd api && \
	buf generate -o sastoj
.PHONY: errors
# generate error proto
errors:
	@cd api && \
	buf generate --template buf.gen.errors.yaml
.PHONY: validate
# generate validate proto
validate:
	@cd api && \
	buf generate --template buf.gen.validate.yaml
.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;

.PHONY: db
# generate db
db:
	go generate ./ent/

.PHONY: docker
# build docker image
docker:
	$(foreach T, $(PROJECTS), sudo docker build --target $(T) -t sastoj/$(T):$(VERSION) . ;)

.PHONY: run
# run a specific service: make run admin [CONF=path/to/config]
run:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make run <service-name> [CONF=path/to/config]"; \
		echo "Available services: admin, contest, auth, gojudge, freshcup"; \
		exit 1; \
	fi

	@SERVICE=$(word 2,$(MAKECMDGOALS)); \
	CONF_PATH=$${CONF:-conf}; \
	case "$$SERVICE" in \
		admin|contest|auth|gojudge|freshcup) \
			echo "Building and running $$SERVICE service..."; \
			case "$$SERVICE" in \
				admin) SVC_PATH="admin/admin" ;; \
				contest) SVC_PATH="user/contest" ;; \
				auth) SVC_PATH="public/auth" ;; \
				gojudge) SVC_PATH="judge/gojudge" ;; \
				freshcup) SVC_PATH="judge/freshcup" ;; \
			esac; \
			go build -o bin/$$SERVICE ./app/$$SVC_PATH/cmd/$$SERVICE; \
			./bin/$$SERVICE -conf $$CONF_PATH/$$SVC_PATH/config.yaml; \
			;; \
		*) \
			echo "Error: Unknown service '$$SERVICE'"; \
			echo "Available services: admin, contest, auth, gojudge, freshcup"; \
			exit 1; \
			;; \
	esac

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
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
