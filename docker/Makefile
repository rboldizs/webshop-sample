DOCKER_DIR=$(CURDIR)
PROJECT_DIR=$(CURDIR)/..

ifndef BINARY_NAME
BINARY_NAME=$(CI_PROJECT_NAME)
endif

BINARY_PATH=$(PROJECT_DIR)/build/$(BINARY_NAME)


ifndef IMAGENAME
IMAGENAME=$(CI_PROJECT_NAME)
endif

ifndef RELEASE_TAG
RELEASE_TAG=latest
endif

clean:
	@echo "clean ..."
	@rm -rf .build
	@echo "done"

prepare:
	@echo "prepare docker build..."
	@mkdir -p .build
	@echo $(BINARY_PATH)
	@cp $(BINARY_PATH) $(DOCKER_DIR)/.build/app
	@cp $(PROJECT_DIR)/server/config/*.json $(DOCKER_DIR)/.build/
	
build: prepare
	@echo "create docker image..."
	@docker build --tag $(IMAGENAME):latest .

push-dev:
	docker tag $(IMAGENAME):latest $(DEV_REGISTRY)/$(DEV_NAMESPACE)/$(IMAGENAME):$(RELEASE_TAG)
	docker tag $(IMAGENAME):latest $(DEV_REGISTRY)/$(DEV_NAMESPACE)/$(IMAGENAME):latest
	@docker login -u none -p $(DEV_REGISTRY_TOKEN) $(DEV_REGISTRY)
	@docker push $(DEV_REGISTRY)/$(DEV_NAMESPACE)/$(IMAGENAME):$(RELEASE_TAG)
	@docker push $(DEV_REGISTRY)/$(DEV_NAMESPACE)/$(IMAGENAME):latest

push-prod:
	@docker tag $(IMAGENAME):latest $(PROD_REGISTRY)/$(PROD_NAMESPACE)/$(IMAGENAME):$(RELEASE_TAG)
	@docker login -u none -p $(PROD_REGISTRY_TOKEN) $(PROD_REGISTRY)
	@docker push $(PROD_REGISTRY)/$(PROD_NAMESPACE)/$(IMAGENAME):$(RELEASE_TAG)

.PHONY: 
