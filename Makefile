APP_NAME=irent-backend
VERSION=latest
PROJECT_ID=sean-side
NS=side
DEPLOY_TO=uat
REGISTRY=gcr.io
IMAGE_NAME=$(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)
HELM_REPO_NAME = blackhorseya

check_defined = $(if $(value $1),,$(error Undefined $1))

.PHONY: help
help: ## show help
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## remove artifacts
	@rm -rf bin coverage.txt profile.out
	@echo Successfully removed artifacts

.PHONY: lint
lint: ## run golint
	@golint -set_exit_status ./...

.PHONY: report
report: ## run report
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/irent'

.PHONY: test-unit
test-unit: ## run unit test
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: build-image
build-image: ## build application image
	$(call check_defined,VERSION)
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--pull --cache-from=$(IMAGE_NAME):latest \
	-f Dockerfile .

.PHONY: list-images
list-images: ## list all images
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images: ## prune image in local
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image
push-image: ## push image to registry
	$(call check_defined,VERSION)
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

.PHONY: deploy
deploy: ## deploy application
	$(call check_defined,VERSION)
	$(call check_defined,DEPLOY_TO)
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) $(HELM_REPO_NAME)/$(APP_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml \
	--set image.tag=$(VERSION)

.PHONY: gen
gen: gen-wire gen-pb gen-swagger gen-mocks ## generate code

.PHONY: gen-wire
gen-wire: ## generate wire code
	@wire gen ./...

.PHONY: gen-pb
gen-pb: ## generate protobuf
	@protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-swagger
gen-swagger: ## generate swagger spec
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal -o api/docs

.PHONY: gen-mocks
gen-mocks: ## generate mocks code via mockery
	@go generate -x ./...
