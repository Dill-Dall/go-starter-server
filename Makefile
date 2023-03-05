.PHONY: openapi
openapi-codegen:
	oapi-codegen -generate types -o ./petapi/openapi_types.gen.go -package petapi todo.yml
	oapi-codegen -generate chi-server -o ./petapi/openapi_server.gen.go -package petapi todo.yml
	oapi-codegen -generate spec -o ./petapi/openapi_spec.go -package petapi todo.yml
	go mod tidy
build:
	go build