
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
	mockgen -source ./services/logs/interface.go -mock_names Service=MockLogService -destination ./tests/mock/log_service.go -package mock
	mockgen -source ./api/middleware/authn/jwt.go -destination ./tests/mock/jwt_provider.go -package mock
	mockgen -source ./pkg/radixapi/client/application/application_client.go -mock_names ClientService=MockRadixApiApplicationClient -destination ./tests/mock/application_client.go -package mock