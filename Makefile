.PHONY: build run start
build:
		go build -v ./cmd/raptor-user-service
run:
	    go run ./cmd/raptor-user-service -env=local -addr=0.0.0.0 -port:8081
start: build run
		
		
.DEFAULT_GOAL := build