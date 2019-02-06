test:
	go test ./...
run:
	go run main.go
swagger:
	swagger generate spec -o swaggerui/swagger.json

