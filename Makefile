include .env

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: clean build

build:
	$(GOBUILD) -o "$(BINARY_DIR)/$(BINARY_NAME)"

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)/

deps:
	$(GOGET) github.com/gorilla/mux
	$(GOGET) github.com/jinzhu/gorm
	$(GOGET) github.com/joho/godotenv
	$(GOGET) github.com/rs/cors
	$(GOGET) github.com/satori/go.uuid
	$(GOGET) github.com/mattn/go-sqlite3
	$(GOGET) github.com/xeipuuv/gojsonschema
	$(GOGET) golang.org/x/time/rate
