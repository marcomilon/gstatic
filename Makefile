GOCMD=go
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

RELEASEDIR=release

all: release
	
setup:
	mkdir -p $(RELEASEDIR)

clean:
	$(GOCLEAN)
	rm -rf $(RELEASEDIR)
	rm -rf website
	
test:
	go test -cover ./...

build-macos: 
	GOOS=darwin $(GOBUILD) -ldflags="-s -w" -o main.go -o $(RELEASEDIR)/gstatic_macos

build-linux:
	GOOS=linux $(GOBUILD) -ldflags="-s -w" -o main.go -o $(RELEASEDIR)/gstatic_linux
	
build-win: setup
	GOOS=windows GOARCH=386 $(GOBUILD) -ldflags="-s -w" -o main.go -o $(RELEASEDIR)/gstatic.exe
	
release: clean build-linux build-win build-macos