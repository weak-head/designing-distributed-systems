IMAGE_ID ?= github.com/weak-head/gomem
IMAGE_TAG ?= 0.0.1

.PHONY: build-docker
build-docker:
	docker build -t $(IMAGE_ID):$(IMAGE_TAG) .

.PHONY: load-image
load-image: build-docker
	kind load docker-image $(IMAGE_ID):$(IMAGE_TAG)