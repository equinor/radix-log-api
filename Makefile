SHELL = bash
.DEFAULT_GOAL = build

.PHONY: swag_tool
swag_tool:
ifeq (, $(shell which swag))
	go install github.com/swaggo/swag/cmd/swag@latest
endif
SWAG_TOOL=$(shell which swag)

.PHONY: staticcheck_tool
staticcheck_tool:
ifeq (, $(shell which staticcheck))
	go install honnef.co/go/tools/cmd/staticcheck@latest
endif
STATICCHECK_TOOL=$(shell which staticcheck)

.PHONY: swagger
swagger: swag_tool
	${SWAG_TOOL} init

.PHONY: test
test:	
	go test -cover ./...

.PHONY: staticcheck
staticcheck: staticcheck_tool
	${STATICCHECK_TOOL} ./...

.PHONY: build
build: swagger
	CGO_ENABLED=0 \
	go build \
	-installsuffix cgo \
	-ldflags="-s -w" \
	-o ./bin/radix-log-api \
	.

.PHONY: radixapiclient
radixapiclient:
	swagger generate client -t ./pkg/radixapi -f https://api.radix.equinor.com/swaggerui/swagger.json -A RadixAPI

.PHONY: mocks
mocks:
	mockgen -source ./api/controllers/interface.go -destination ./api/controllers/mock_controller.go -package controllers
	mockgen -source ./services/logs/interface.go -mock_names Service=MockLogService -destination ./services/logs/mock_service.go -package logs
	mockgen -source ./api/middleware/authn/jwt_validator.go -destination ./api/middleware/authn/mock_jwt_validator.go -package authn
	mockgen -source ./api/middleware/authn/provider.go -destination ./api/middleware/authn/mock_provider.go -package authn
	mockgen -source ./api/middleware/authz/policy.go -destination ./api/middleware/authz/mock_policy.go -package authz
	mockgen -source ./pkg/radixapi/client/application/application_client.go -mock_names ClientService=MockRadixApiApplicationClient -destination ./internal/tests/mock/application_client.go -package mock