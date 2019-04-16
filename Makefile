BINARY_DIR=$(CURDIR)/build
BINARY_PATH=$(BINARY_DIR)/$(CI_PROJECT_NAME)

TEMPLATE_DIR=$(CURDIR)/shop-backend/templates/api/
DOCKER_DIR=$(CURDIR)/docker

export BINARY_PATH
export DOCKER_DIR

VERSION=1.0.0
BUILD_TIME=`date +%FT%T%Z`

ifndef $(CI_COMMIT_SHA)
CI_COMMIT_SHA="commit-not-set"
endif

LDFLAGS=-ldflags ""

generate-build-info:
	@echo "generating application build info ..."
	@sed -e "s/##VERSION##/${VERSION}/g" \
		-e "s/##BUILD_ID##/${CI_BUILD_ID}/g" \
		-e "s/##BUILD_TIME##/${BUILD_TIME}/g" \
		-e "s/##GIT_COMMIT##/${CI_COMMIT_SHA}/g" \
		$(TEMPLATE_DIR)/status.html 

build: generate-build-info
	@echo "start building ..."

test:
	@echo "test ..."
	@$(GOTEST_UNIT) -cover ./...

checkstyle:
	@echo "stylecheck ..."
	@golint `go list ./...`

