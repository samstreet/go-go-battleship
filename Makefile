include .env

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: clean build

build:
	$(GOBUILD) -o $(BINARY_NAME)
clean:
	$(GOCLEAN)
	rm -rf bin/
deps:
	$(GOGET) github.com/gorilla/mux
	$(GOGET) github.com/jinzhu/gorm
	$(GOGET) github.com/joho/godotenv