#Go parameters

GOCMD=Go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) GOGET
BINARY_NAME=gocopy


all: test build

build:
      $(GOBUILD) -o $(BINARY_NAME) -v
test:
      $(GOTEST) -v ./..
clean: 
      $(GOCLEAN) 
	  rm -f $(BINARY_NAME)
run: 
    $(GOBUILD) -o $(BINARY_NAME) -v ./..
	./$(BINARY_NAME)
deps:
    $(GOGET) github.com/cheggaaa/pb/v3



# cross compile

build-linux: 
            GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
build-windows:
            GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

 