BINARY_DIR=$(CURDIR)/build
BINARY_PATH=$(BINARY_DIR)/$(CI_PROJECT_NAME)
MOCK_DIR=$(CURDIR)/mocks
INFO_DIR=$(CURDIR)/info
TEMPLATE_DIR=$(CURDIR)/template
DOCKER_DIR=$(CURDIR)/docker

export BINARY_PATH
export DOCKER_DIR

GO=go
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean -x
GOBUILD=$(GO) build
GOGENERATE=$(GO) generate
GOTEST_INTEGRATION=$(GO) test -test.v
GOTEST_UNIT=$(GO) test -test.v -test.short
GOGET=$(GO) get -u
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%Z`

ifndef $(CI_COMMIT_SHA)
CI_COMMIT_SHA="commit-not-set"
endif

LDFLAGS=-ldflags ""

$(eval $(call golang-version-check,1.6))

go-version:
	@$(GO) version
	@echo "make build|clean|test|install|ci"

clean:
	@echo "clean ..."
	@$(GOCLEAN)
	@if [ -d $(BINARY_DIR) ] ; then rm -rf $(BINARY_DIR); fi
	@if [ -d $(INFO_DIR) ] ; then rm -rf $(INFO_DIR); fi
	@if [ -d vendor ] ; then rm -rf vendor; fi

test-mockgen: generate-build-info
	@echo "generating mocks ..."
	@if [ ! -d $(MOCK_DIR) ] ; then mkdir -p $(MOCK_DIR); fi
	@$(GOGENERATE) ./...

generate-build-info:
	@echo "generating application infos ..."
	@if [ ! -d $(INFO_DIR) ] ; then mkdir -p $(INFO_DIR); fi
	@sed -e "s/##VERSION##/${VERSION}/g" \
		-e "s/##BUILD_ID##/${CI_BUILD_ID}/g" \
		-e "s/##BUILD_TIME##/${BUILD_TIME}/g" \
		-e "s/##GIT_COMMIT##/${CI_COMMIT_SHA}/g" \
		$(TEMPLATE_DIR)/application-info.go.template > $(INFO_DIR)/application-info.go

build: generate-build-info
	@echo "start building ..."
	@(cd server; $(GOBUILD) -o $(BINARY_PATH) $(LDFLAGS))

test:
	@echo "test ..."
	@$(GOTEST_UNIT) -cover ./...

checkstyle:
	@echo "stylecheck ..."
	@golint `go list ./...`

get-prerequisites:
	@echo "get packages ..."
	@# Install as documented: https://github.com/golang/mock
	@$(GOGET) github.com/golang/mock/gomock
	@$(GOINSTALL) github.com/golang/mock/mockgen
	@$(GOGET) github.com/golang/dep/cmd/dep
	@$(GOGET) golang.org/x/lint/golint

start-server: build
	@(cd $(BINARY_DIR); $(BINARY_PATH) -v=2 -logtostderr=true)
	
	
get-dependencies:
	@echo "get dependencies..."
	@dep ensure -v


dockerimage:
	@echo "create docker image..."
	$(DOCKER_DIR)/build.sh

install: get-prerequisites get-dependencies test-mockgen

ci: install integration-test checkstyle generate-build-info build

.PHONY: all test clean build