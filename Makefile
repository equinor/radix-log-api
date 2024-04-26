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
	AZURE_CLIENT_ID=eef3ea0a-a1f8-4eb2-bac7-1c63b939e457 LOG_API_LOG_ANALYTICS_WORKSPACE_ID=744f4b25-b47d-4ad2-8f5d-f97382a81251 envsubst < radixconfig.tpl.yaml > radixconfig.dev.yaml
	AZURE_CLIENT_ID=d5ce79f3-3b26-46ad-992c-59f9a913b635 LOG_API_LOG_ANALYTICS_WORKSPACE_ID=7dffa030-c9ac-4d34-90ed-77ee26511aa4 envsubst < radixconfig.tpl.yaml > radixconfig.playground.yaml
	AZURE_CLIENT_ID=a9489564-b5aa-404e-8a82-a21f2faa175f LOG_API_LOG_ANALYTICS_WORKSPACE_ID=ae599795-733e-48d5-b4bf-141c5e10f6b1 envsubst < radixconfig.tpl.yaml > radixconfig.platform.yaml
	AZURE_CLIENT_ID=ac2eb093-8396-4f37-b43c-4bef2b5f9948 LOG_API_LOG_ANALYTICS_WORKSPACE_ID=4f4ff917-823c-4f80-8a46-4fbe7de5a32e envsubst < radixconfig.tpl.yaml > radixconfig.c2.yaml

.PHONY: generate
generate: radixconfigs mocks swagger

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
