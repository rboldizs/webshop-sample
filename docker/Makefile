DOCKER_DIR=$(CURDIR)
PROJECT_DIR=$(CURDIR)/..
TEMPLATE_DIR=$(CURDIR)/.build/shop_backend/templates/api/

ifndef BINARY_NAME
BINARY_NAME=$(CI_PROJECT_NAME)
endif

BINARY_PATH=$(PROJECT_DIR)/shop_backend


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

generate-build-info: prepare
	@echo "generating application build info ..."
	@sed -e "s/##VERSION##/${VERSION}/g" \
		-e "s/##BUILD_ID##/${CI_BUILD_ID}/g" \
		-e "s/##BUILD_TIME##/${BUILD_TIME}/g" \
		-e "s/##GIT_COMMIT##/${CI_COMMIT_SHA}/g" \
		$(TEMPLATE_DIR)/status.html 

prepare:
	@echo "prepare docker build..."
	@mkdir -p .build
	@cp -rf $(BINARY_PATH) $(DOCKER_DIR)/.build/
	
build: generate-build-info 
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
