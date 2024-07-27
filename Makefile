check_install:
	@which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	@swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	@cd sdk && swagger generate client -f ../swagger.yaml -A [BOOKNEST-API]

build: swagger
	@go build -o ./bin/booknest

run: build
	@./bin/booknest

test_store:
	@go test ./store -v

test: generate_client
	@go test -v