.PHONY: build run

SERVICE = bin/wx-project
MAIN = cmd/main

default: build 

build:
	go1.16 build -o ./$(SERVICE) ./$(MAIN)

run: build
	./$(SERVICE)