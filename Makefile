build:
	@go build -o ./bin/booknest
run: build
	@./bin/booknest
test:
	@go test ./store -v