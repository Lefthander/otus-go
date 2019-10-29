#Go parameters

GOCMD=Go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

BINARY_NAME=goenvdir


all: build

build:
      $(GOBUILD) -o $(BINARY_NAME) -v
clean: 
      $(GOCLEAN) 
	  rm -f $(BINARY_NAME)
run: 
    $(GOBUILD) -o $(BINARY_NAME) -v ./..
	./$(BINARY_NAME)


# cross compile

build-linux: 
            GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
build-windows:
            GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

 