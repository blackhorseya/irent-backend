APP_NAME=irent-backend
VERSION=latest
PROJECT_ID=sean-side
NS=side
DEPLOY_TO=uat
REGISTRY=gcr.io
IMAGE_NAME=$(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)
HELM_REPO_NAME = blackhorseya

check_defined = $(if $(value $1),,$(error Undefined $1))

.PHONY: clean
clean:
	@rm -rf bin coverage.txt profile.out
	@echo Successfully removed artifacts

.PHONY: lint
lint:
	@golint -set_exit_status ./...

.PHONY: report
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/irent'

.PHONY: test-unit
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: build-image
build-image:
	$(call check_defined,VERSION)
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--pull --cache-from=$(IMAGE_NAME):latest \
	-f Dockerfile .

.PHONY: list-images
list-images:
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: push-image
push-image:
	$(call check_defined,VERSION)
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

.PHONY: install-db
install-db:
	@helm --namespace $(NS) upgrade --install $(APP_NAME)-db bitnami/redis \
	--values ./deployments/configs/$(DEPLOY_TO)/redis.yaml

.PHONY: deploy
deploy:
	$(call check_defined,VERSION)
	$(call check_defined,DEPLOY_TO)
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) $(HELM_REPO_NAME)/$(APP_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml \
	--set image.tag=$(VERSION)

.PHONY: gen
gen: gen-wire gen-pb gen-swagger

.PHONY: gen-wire
gen-wire:
	@wire gen ./...

.PHONY: gen-pb
gen-pb:
	@protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-swagger
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal -o api/docs
