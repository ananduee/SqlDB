test:
	go test 

build:
	go build -o SqlDB main.go 

.DEFAULT_GOAL := build