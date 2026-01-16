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

# Allow module names as targets without defining explicit rules for them
.PHONY: admin contest auth cases gojudge freshcup gojudge-server user-contest public-auth
admin contest auth cases gojudge freshcup gojudge-server user-contest public-auth:
	@:

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
# generate protobuf api go code: make api [module]
api:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then \
		cd api && \
		buf generate -o sastoj; \
	else \
		MODULE=$(word 2,$(MAKECMDGOALS)); \
		case "$$MODULE" in \
			admin) \
				echo "Generating API code for admin module..."; \
				cd api && buf generate -o sastoj sastoj/admin/admin/service/v1/admin.proto; \
				;; \
			contest) \
				echo "Generating API code for contest module..."; \
				cd api && buf generate -o sastoj sastoj/user/contest/service/v1/contest.proto; \
				;; \
			auth) \
				echo "Generating API code for auth module..."; \
				cd api && buf generate -o sastoj sastoj/public/auth/service/v1/auth.proto; \
				;; \
			cases) \
				echo "Generating API code for cases module..."; \
				cd api && buf generate -o sastoj sastoj/rsjudge/cases/service/v1/cases.proto; \
				;; \
			*) \
				echo "Error: Unknown module '$$MODULE'"; \
				echo "Available modules: admin, contest, auth, cases"; \
				exit 1; \
				;; \
		esac; \
	fi

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
# build docker image: make docker [target]
docker:
	@if [ -z "$(word 2,$(MAKECMDGOALS))" ]; then \
		$(foreach T, $(PROJECTS), docker build --target $(T) -t sastoj/$(T):$(VERSION) . ;) \
	else \
		TARGET=$(word 2,$(MAKECMDGOALS)); \
		docker build --target $$TARGET -t sastoj/$$TARGET:$(VERSION) . ; \
	fi

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
			case "$$SERVICE" in \
				gojudge) \
					./bin/$$SERVICE -conf $$CONF_PATH/gojudge/config.yaml; \
					;; \
				freshcup) \
					./bin/$$SERVICE -conf $$CONF_PATH/freshcup/config.yaml; \
					;; \
				*) \
					./bin/$$SERVICE -conf $$CONF_PATH/$$SVC_PATH/config.yaml; \
					;; \
			esac; \
			;; \
		*) \
			echo "Error: Unknown service '$$SERVICE'"; \
			echo "Available services: admin, contest, auth, gojudge, freshcup"; \
			exit 1; \
			;; \
	esac

.PHONY: run-sandbox
# run standalone go-judge sandbox for local debugging
run-sandbox:
	docker run -d \
		--name go-judge-standalone \
		--privileged \
		--shm-size=256m \
		-p 5050:5050 \
		-p 5051:5051 \
		-v $(shell pwd)/data:/data \
		criyle/go-judge:v1.11.3 \
		-enable-grpc

.PHONY: init-sandbox
# install toolchain in the standalone sandbox
init-sandbox:
	docker exec -i go-judge-standalone sh -c "cat > /tmp/init_sandbox.sh" < scripts/init_sandbox.sh
	docker exec go-judge-standalone chmod +x /tmp/init_sandbox.sh
	docker exec go-judge-standalone /tmp/init_sandbox.sh

.PHONY: help
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
	@echo ''
	@echo 'Module-specific commands:'
	@echo ' make api <module>      Generate API code for a specific module'
	@echo ' make errors <module>   Generate error code for a specific module'
	@echo ' make validate <module> Generate validate code for a specific module'
	@echo ''
	@echo 'Available modules: admin, contest, auth, cases'
	@echo 'Available services: admin, contest, auth, gojudge, freshcup'

.DEFAULT_GOAL := help
