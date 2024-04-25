SHELL = bash
.DEFAULT_GOAL = build


.PHONY: swagger
swagger: bootstrap
	swag init

.PHONY: test
test:
	go test -cover ./...

.PHONY: lint
lint: bootstrap
	golangci-lint run --max-same-issues 0

.PHONY: build
build:
	CGO_ENABLED=0 \
	go build \
	-installsuffix cgo \
	-ldflags="-s -w" \
	-o ./bin/radix-log-api \
	.

.PHONY: radixapiclient
radixapiclient: bootstrap
	swagger generate client -t ./pkg/radixapi -f https://api.radix.equinor.com/swaggerui/swagger.json -A RadixAPI

.PHONY: mocks
mocks: bootstrap
	mockgen -source ./api/controllers/interface.go -destination ./api/controllers/mock_controller.go -package controllers
	mockgen -source ./pkg/services/logs/interface.go -mock_names Service=MockLogService -destination ./pkg/services/logs/mock_service.go -package logs
	mockgen -source ./api/middleware/authn/jwt_validator.go -destination ./api/middleware/authn/mock_jwt_validator.go -package authn
	mockgen -source ./api/middleware/authn/provider.go -destination ./api/middleware/authn/mock_provider.go -package authn
	mockgen -source ./api/middleware/authz/policy.go -destination ./api/middleware/authz/mock_policy.go -package authz
	mockgen -source ./pkg/radixapi/client/application/application_client.go -mock_names ClientService=MockRadixApiApplicationClient -destination ./internal/tests/mock/application_client.go -package mock

.PHONY: radixconfigs
radixconfigs:
	AZURE_CLIENT_ID=31816177-d769-4392-8430-b2357f59c701 envsubst < radixconfig.yaml.tpl > radixconfig.dev.yaml
	AZURE_CLIENT_ID=64ead1ac-43da-4fd6-b1c1-6c6e9747dedc envsubst < radixconfig.yaml.tpl > radixconfig.playground.yaml
	AZURE_CLIENT_ID=4ce6649f-1e7d-4293-8ecf-411ddfa00dce envsubst < radixconfig.yaml.tpl > radixconfig.platform.yaml
	AZURE_CLIENT_ID=a8b35d63-7baa-4de1-8d13-ee2d49c6c944 envsubst < radixconfig.yaml.tpl > radixconfig.c2.yaml

.PHONY: generate
generate: generate-radixconfig # radixapiclient mocks swagger

.PHONY: verify-generate
verify-generate: generate
	git diff --exit-code

HAS_SWAG_TOOL     := $(shell command -v swag)
HAS_SWAGGER       := $(shell command -v swagger;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
HAS_MOCKGEN       := $(shell command -v mockgen;)

bootstrap:
ifndef HAS_SWAGGER
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
endif
ifndef HAS_GOLANGCI_LINT
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
endif
ifndef HAS_MOCKGEN
	go install github.com/golang/mock/mockgen@v1.6.0
endif
ifndef HAS_SWAG_TOOL
	go install github.com/swaggo/swag/cmd/swag@latest
endif
