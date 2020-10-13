# Scrapper package
SCRAPER_PACKAGE_FILE = scraper_package.zip

# Web container
WEB_IMAGE_ID ?= github.com/weak-head/faas-web
WEB_IMAGE_TAG ?= 0.0.1

# App container
APP_IMAGE_ID ?= github.com/weak-head/faas-app
APP_IMAGE_TAG ?= 0.0.1


.PHONY: build-docker
build-docker:
	# Build Nginx + Frontend image
	cd web && docker build -t $(WEB_IMAGE_ID):$(WEB_IMAGE_TAG) .

	# Build backend image
	cd app && docker build -t $(APP_IMAGE_ID):$(APP_IMAGE_TAG) .

.PHONY: load-image
load-image: build-docker
	# Load Nginx + Frontend image into kind
	kind load docker-image $(WEB_IMAGE_ID):$(WEB_IMAGE_TAG)

	# Load backend image into kind
	kind load docker-image $(APP_IMAGE_ID):$(APP_IMAGE_TAG)

.PHONY: scraper-package
scraper-package:
	# Install scraper dependencies locally
	cd scraper && pip3 install -t . bs4 requests 

	# Package function with dependencies
	cd scraper && zip -r9 ./${SCRAPER_PACKAGE_FILE} *

.PHONY: clean
clean:
	# Delete scraper function
	kubeless function delete scraper

	kubectl delete deployment faas-deployment
	kubectl delete service faas-service

.PHONY: deploy
deploy:
	# Deploy scrapper function
	kubectl apply -f scraper.yaml

	# kubectl get functions
	# kubeless function ls
	# kubeless function call scraper --data '{"url":"https://webscraper.io/test-sites/e-commerce/allinone"}'

	kubectl apply -f webapp-deployment.yaml
	kubectl apply -f webapp-service.yaml

	# kubectl port-forward service/faas-service 8888:80
	# curl http://localhost:8888