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

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%Z`

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
		-i $(TEMPLATE_DIR)/status.html 

prepare:
	@echo "prepare docker build..."
	@mkdir -p .build
	@cp -rf $(BINARY_PATH) $(DOCKER_DIR)/.build/
	
build: generate-build-info 
	@echo "create docker image..."
	@docker build --tag $(IMAGENAME):latest .

docker: build
	docker tag $(IMAGENAME):latest $(MCPIPE_REGISTRY_SERVER)/$(IMAGENAME):latest
	@docker login -u $(MCPIPE_REGISTRY_USER) -p $(MCPIPE_REGISTRY_TOKEN) $(MCPIPE_REGISTRY_SERVER)
	@docker push $(MCPIPE_REGISTRY_SERVER)/$(IMAGENAME):latest


.PHONY: 
