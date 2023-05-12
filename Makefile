
.PHONY: test
test:
	go test -cover `go list ./...`

.PHONY: swagger
swagger:
	swag init


.PHONY: radixapiclient
radixapiclient:
	swagger generate client -t ./pkg/radixapi -f https://api.radix.equinor.com/swaggerui/swagger.json -A RadixAPI

.PHONY: mocks
mocks:
	mockgen -source ./api/controllers/interface.go -destination ./api/controllers/mock_controller.go -package controllers
	mockgen -source ./services/logs/interface.go -mock_names Service=MockLogService -destination ./services/logs/mock_service.go -package logs
	mockgen -source ./api/middleware/authn/jwt_validator.go -destination ./api/middleware/authn/mock_jwt_validator.go -package authn
	mockgen -source ./api/middleware/authn/provider.go -destination ./api/middleware/authn/mock_provider.go -package authn
	mockgen -source ./pkg/radixapi/client/application/application_client.go -mock_names ClientService=MockRadixApiApplicationClient -destination ./internal/tests/mock/application_client.go -package mock