GO				?= go
SUDO			?= sudo
CONTAINER_TOOL	?= podman
PROJECT			?= mai-backend

DOCKERFILES_DIR			:= deploy/docker/
API_SERVER_DOCKERFILE	:= $(DOCKERFILES_DIR)/api-server.dockerfile
UI_SERVER_DOCKERFILE	:= $(DOCKERFILES_DIR)/ui-server.dockerfile

API_SERVER_IMAGE_NAME		?= api-server
API_SERVER_IMAGE_VERSION	?= latest
API_SERVER_IMAGE			:= $(API_SERVER_IMAGE_NAME):$(API_SERVER_IMAGE_VERSION)

UI_SERVER_IMAGE_NAME	?= ui-server
UI_SERVER_IMAGE_VERSION	?= latest
UI_SERVER_IMAGE			:= $(UI_SERVER_IMAGE_NAME):$(UI_SERVER_IMAGE_VERSION)

# Build dirs

BUILD_CACHE_DIR := ./.build-cache

BUILD_DIRS := $(BUILD_CACHE_DIR)
$(BUILD_DIRS):
	mkdir -p $(BUILD_DIRS)

# Build

define build-image =
	$(CONTAINER_TOOL) build -t $(TAG) -f $(DOCKERFILE) .
endef

.PHONY: build-image-api-server
build-image-api-server: TAG = $(API_SERVER_IMAGE)
build-image-api-server: DOCKERFILE = $(API_SERVER_DOCKERFILE)
build-image-api-server:
	$(build-image)

.PHONY: build-image-ui-server
build-image-ui-server: TAG = $(UI_SERVER_IMAGE)
build-image-ui-server: DOCKERFILE = $(UI_SERVER_DOCKERFILE)
build-image-ui-server:
	$(build-image)

.PHONY: build-images
build-images: \
	build-image-api-server \
	build-image-ui-server

# Start

POD						?= $(PROJECT)
API_SERVER_CONTAINER	?= $(POD)-api-server
UI_SERVER_CONTAINER		?= $(POD)-ui-server

ifeq ($(CONTAINER_TOOL),podman)
	CONTAINER_ENCLOSURE ?= "pod"
	API_SERVER_PORTS	:=
	UI_SERVER_PORTS		:=
	UI_SERVER_ENV		:= --env API_SERVER_ADDR=http://$(POD):8080
else
	CONTAINER_ENCLOSURE ?= "network"
	API_SERVER_PORTS	:= -p 8080:8080
	UI_SERVER_PORTS		:= -p 80:80
	UI_SERVER_ENV		:= --env API_SERVER_ADDR=http://$(API_SERVER_CONTAINER):8080
endif

.PHONY: container-env
container-env:
ifeq ($(CONTAINER_TOOL),podman)
	-$(CONTAINER_TOOL) network create $(POD)
	$(CONTAINER_TOOL) $(CONTAINER_ENCLOSURE) create --replace \
		--name $(POD) \
		--network $(POD) \
		-p 8080:8080 \
		-p 80:80
else
	-$(CONTAINER_TOOL) $(CONTAINER_ENCLOSURE) create $(POD)
endif

define start-container =
	$(CONTAINER_TOOL) run --rm -d --$(CONTAINER_ENCLOSURE) $(POD)
endef

define stop-container =
	-@$(CONTAINER_TOOL) stop
endef

.PHONY: stop-ui-server
stop-ui-server:
	$(stop-container) $(UI_SERVER_CONTAINER)

.PHONY: start-ui-server
start-ui-server: stop-ui-server
	$(start-container) \
		--name $(UI_SERVER_CONTAINER) \
		$(UI_SERVER_ENV) \
		$(UI_SERVER_PORTS) \
		$(UI_SERVER_IMAGE)

.PHONY: stop-api-server
stop-api-server:
	$(stop-container) $(API_SERVER_CONTAINER)

.PHONY: start-api-server
start-api-server: stop-api-server
	$(start-container) \
		--name $(API_SERVER_CONTAINER) \
		$(API_SERVER_PORTS) \
		$(API_SERVER_IMAGE)

# All

.PHONY: build
build: build-images

.PHONY: stop
stop: \
	stop-api-server \
	stop-ui-server

.PHONY: start
start: container-env \
	start-api-server \
	start-ui-server

# Docker compose

.PHONY: up
up:
	docker-compose up -d --remove-orphans

.PHONY: down
down:
	docker-compose down

.PHONY: build
	docker-compose build
