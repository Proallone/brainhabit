# Detect the OS
.PHONY: deps, run, build, test, compose-up, compose-down, init, docs, check-os clean
NAME=bh-backend
API_MAIN=cmd/api/main.go

## deps: Download modules
deps:
	@go mod download
## run: Build and Run
run : build
	@./bin/$(NAME)

## build: Build and Run in development mode.
build : 
	@go build -o bin/$(NAME) $(API_MAIN) 

## test: Run tests with verbose mode
test : 
	@go test -v ./tests

## compose-up: Build from Dockerfile, pull other images and start all containers
compose-up : 
	@docker-compose up -d --no-deps --build

## compose-down: clear all containers and volumes
compose-down : 
	@docker-compose down --volumes
