# DO NOT MODIFY

# ################################
# Load configuration
SHELL := /bin/bash

# If there is a dependency proxy configured in GitLab, let's use it, otherwise fallback to docker.io
DOCKER_PREFIX=
ifneq (${CI_PIPELINE_ID},) # If we're using the GitLab CI/CD pipeline
DOCKER_PREFIX=git.sapienzaapps.it:443/lab/dependency_proxy/containers/
endif
export DOCKER_PREFIX

OPENAPI_CLI="${DOCKER_PREFIX}openapitools/openapi-generator-cli:v5.4.0"
SWAGGER_UI="${DOCKER_PREFIX}swaggerapi/swagger-ui:v4.11.1"

# Common variables
GITORIGIN := $(shell git remote get-url origin | sed 's/.git$$//' | sed 's/git@git.sapienzaapps.it://' | sed 's$$https://.*git.sapienzaapps.it/$$$$')
PROJECT_NAME := $(shell echo ${GITORIGIN} | cut -d '/' -f 2 | tr '[:upper:]' '[:lower:]')
GROUP_NAME := $(shell echo ${GITORIGIN} | cut -d '/' -f 1 | tr '[:upper:]' '[:lower:]')
VERSION := $(shell git fetch --unshallow >/dev/null 2>&1; git describe --all --long --dirty 2>/dev/null)
ifeq (${VERSION},)
VERSION := no-git-version
endif

DOCKER_IMAGE_PATH := registry.git.sapienzaapps.it/${GROUP_NAME}/${PROJECT_NAME}
ifneq (${CI_REGISTRY_IMAGE},) # If we're using the GitLab CI/CD pipeline
DOCKER_IMAGE_PATH := ${CI_REGISTRY_IMAGE}
endif

BUILD_ID := ${CI_PIPELINE_ID}
ifeq (${BUILD_ID},) # Fallback for local builds
BUILD_ID = 0
endif

BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
REPO_HASH := $(shell git rev-parse --verify HEAD)
DOCKER_IMAGE_TAG := $(shell date -u +"%Y-%m-%d").${BUILD_ID}

GO_MODULE := $(shell cat go.mod | grep module | cut -f 2 -d ' ')

# ################################
# Targets

.PHONY: all
all: build

# Build/dev section
.PHONY: info
info:
	$(info ** Environment info **)
	$(info Project name: ${PROJECT_NAME})
	$(info Group name: ${GROUP_NAME})
	$(info Version: ${VERSION})
	$(info Docker image path: ${DOCKER_IMAGE_PATH})
	$(info Docker image tag: ${DOCKER_IMAGE_TAG})

.PHONY: env
env:
	$(info PROJECT_NAME="${PROJECT_NAME}")
	$(info GROUP_NAME="${GROUP_NAME}")
	$(info VERSION="${VERSION}")
	$(info DOCKER_IMAGE_PATH="${DOCKER_IMAGE_PATH}")
	$(info DOCKER_IMAGE_TAG="${DOCKER_IMAGE_TAG}")
	@:

.PHONY: build
build:
	go generate -mod=vendor ./...
	/bin/bash -euo pipefail -c "cd cmd; for ex in \$$(ls); do pushd \$$ex; CGO_ENABLED=0 go build -mod=vendor -ldflags \"-extldflags \\\"-static\\\"\" -a -installsuffix cgo -o ../../\$$ex.exe .; popd; done"

.PHONY: docker
docker: info
	docker build \
		-t ${DOCKER_IMAGE_PATH}:${DOCKER_IMAGE_TAG} \
		--build-arg DOCKER_PREFIX="${DOCKER_PREFIX}" \
		--build-arg PROJECT_NAME="${PROJECT_NAME}" \
		--build-arg GROUP_NAME="${GROUP_NAME}" \
		--build-arg APP_VERSION="${VERSION}" \
		--build-arg BUILD_DATE="${BUILD_DATE}" \
		--build-arg REPO_HASH="${REPO_HASH}" \
		.

.PHONY: push
push:
	docker push ${DOCKER_IMAGE_PATH}:${DOCKER_IMAGE_TAG}
	docker tag ${DOCKER_IMAGE_PATH}:${DOCKER_IMAGE_TAG} ${DOCKER_IMAGE_PATH}:latest
	docker push ${DOCKER_IMAGE_PATH}:latest

.PHONY: up-deps
up-deps:
	docker-compose -p ${PROJECT_NAME} \
		-f demo/docker-compose.yml \
		up

.PHONY: stop
stop:
	DOCKER_IMAGE=${DOCKER_IMAGE_PATH}:${DOCKER_IMAGE_TAG} \
		docker-compose -p ${PROJECT_NAME} \
		-f demo/docker-compose.yml \
		-f demo/docker-compose.cicd.yml \
		stop

.PHONY: logs
logs:
	DOCKER_IMAGE=${DOCKER_IMAGE_PATH}:${DOCKER_IMAGE_TAG} \
		docker-compose -p ${PROJECT_NAME} \
		-f demo/docker-compose.yml \
		-f demo/docker-compose.cicd.yml \
		logs > docker-compose.log

.PHONY: down
down:
	DOCKER_IMAGE=${DOCKER_IMAGE_PATH}:${DOCKER_IMAGE_TAG} \
		docker-compose -p ${PROJECT_NAME} \
		-f demo/docker-compose.yml \
		-f demo/docker-compose.cicd.yml \
		down

.PHONY: install-dev-deps
install-dev-deps:
	go install github.com/psampaz/go-mod-outdated@v0.8.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0
	go install github.com/google/go-licenses@v1.2.1
	go install golang.org/x/tools/cmd/godoc@latest
	go install github.com/amacneil/dbmate@v1.15.0

.PHONY: doc
doc:
	docker run --rm -it \
		--user $(shell id -u):$(shell id -g) \
		-v "$(shell pwd)/doc/:/local/" \
		${OPENAPI_CLI} generate -g openapi -i /local/api.yaml -o /local/json/
	docker run --rm -it \
		--user $(shell id -u):$(shell id -g) \
		-v "$(shell pwd)/doc/:/doc/:ro" \
		-p 9000:8080 \
		-e SWAGGER_JSON=/doc/json/openapi.json \
		${SWAGGER_UI}

.PHONY: godoc
godoc:
	$(info Open http://localhost:6060/pkg/${GO_MODULE}/)
ifdef ($GOROOT,)
	godoc
else
	godoc -goroot /usr/share/go
endif

.PHONY: beta
beta:
ifeq ($(shell git status --porcelain),)
	git pull origin master
	git checkout beta
	git pull origin beta
	git merge --ff origin/master
	git push origin beta
	git checkout master
else
	$(error Please commit all files first)
endif
