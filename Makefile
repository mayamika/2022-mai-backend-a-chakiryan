SHELL			?= /usr/bin/env bash
GO				?= go
SUDO			?= sudo
CONTAINER_TOOL	?= docker
PROJECT			?= mai-backend

DOCKERFILES_DIR			:= deploy/docker/
API_SERVER_DOCKERFILE	:= $(DOCKERFILES_DIR)/api-server.dockerfile
UI_SERVER_DOCKERFILE	:= $(DOCKERFILES_DIR)/ui-server.dockerfile

CONTAINERS_REGISTRY ?= cr.yandex/${YC_CONTAINER_REGISTRY}

API_SERVER_IMAGE_NAME		?= api-server
API_SERVER_IMAGE_VERSION	?= 0.1.0
API_SERVER_IMAGE			:= $(API_SERVER_IMAGE_NAME):$(API_SERVER_IMAGE_VERSION)

UI_SERVER_IMAGE_NAME	?= ui-server
UI_SERVER_IMAGE_VERSION	?= 0.1.0
UI_SERVER_IMAGE			:= $(UI_SERVER_IMAGE_NAME):$(UI_SERVER_IMAGE_VERSION)

# Build dirs

BUILD_CACHE_DIR := ./.build-cache

BUILD_DIRS := $(BUILD_CACHE_DIR)
$(BUILD_DIRS):
	mkdir -p $(BUILD_DIRS)

# Build images

define push-image =
	$(CONTAINER_TOOL) tag $(TAG) $(CONTAINERS_REGISTRY)/$(TAG)
	$(CONTAINER_TOOL) push $(CONTAINERS_REGISTRY)/$(TAG)
endef

define build-image =
	$(CONTAINER_TOOL) build -t $(TAG) -f $(DOCKERFILE) .
endef

.PHONY: build-image-api-server
build-image-api-server: TAG = $(API_SERVER_IMAGE)
build-image-api-server: DOCKERFILE = $(API_SERVER_DOCKERFILE)
build-image-api-server:
	$(build-image)

.PHONY: push-image-api-server
push-image-api-server: TAG = $(API_SERVER_IMAGE)
push-image-api-server:
	$(push-image)

.PHONY: build-image-ui-server
build-image-ui-server: TAG = $(UI_SERVER_IMAGE)
build-image-ui-server: DOCKERFILE = $(UI_SERVER_DOCKERFILE)
build-image-ui-server:
	$(build-image)

.PHONY: push-image-ui-server
push-image-ui-server: TAG = $(UI_SERVER_IMAGE)
push-image-ui-server:
	$(push-image)

.PHONY: build-images
build-images: \
	build-image-api-server \
	build-image-ui-server

.PHONY: push-images
push-images: \
	push-image-api-server \
	push-image-ui-server

# Docker compose

.PHONY: compose-build
	docker-compose build

.PHONY: compose-up
compose-up:
	docker-compose up -d --remove-orphans

.PHONY: compose-down
compose-down:
	docker-compose down
