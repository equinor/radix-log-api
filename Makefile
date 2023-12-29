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
