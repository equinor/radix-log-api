
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
	mockgen -source ./api/controllers/interface.go -destination ./internal/tests/mock/controller.go -package mock
	mockgen -source ./services/logs/interface.go -mock_names Service=MockLogService -destination ./internal/tests/mock/log_service.go -package mock
	mockgen -source ./api/middleware/authn/jwt.go -destination ./internal/tests/mock/jwt_validator.go -package mock
	mockgen -source ./api/middleware/authn/provider.go -destination ./internal/tests/mock/authn_provider.go -package mock
	mockgen -source ./pkg/radixapi/client/application/application_client.go -mock_names ClientService=MockRadixApiApplicationClient -destination ./internal/tests/mock/application_client.go -package mock