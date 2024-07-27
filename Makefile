check_install:
	@which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	@swagger generate spec -o ./swagger.yaml --scan-models

generate_client:
	@mkdir -p sdk
	@cd sdk && swagger generate client -f ../swagger.yaml -A [BOOKNEST-API]

build: swagger
	@go build -o ./bin/booknest

run: build
	@./bin/booknest

css:
	@tailwindcss -i ./views/css/app.css -o ./public/styles.css --watch

testing:
	@air
	@templ generate --watch --proxy=http://localhost:7000

test_store:
	@go test ./store -v

test: generate_client
	@go test -v
