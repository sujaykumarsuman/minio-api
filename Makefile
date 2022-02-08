oapi-code:
	mkdir -p api
	oapi-codegen -package openapi3 -generate server -o server.gen.go openapi.yml
	oapi-codegen -generate spec -package api ./openapi.yml > api/spec.go
	oapi-codegen -generate types -package api ./openapi.yml > api/models.go
	oapi-codegen -generate chi-server -package api ./openapi.yml > api/server.go