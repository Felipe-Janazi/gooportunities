.PHONY: default run build test docs clean
#Variables
APP_NAME=gooportunities

#Tasks
default: run-with-docs

run: 
	@go run main.go
run-with-docs: 
	@swag init
	@go run  main.go
build: 
	@go build -o $(APP_NAME) main_go
test:
	@go test ./ ...
docs: 
	@swag init
clean: 
	@rm -f $(APP_NAME)
	@rm -rf ./docs
