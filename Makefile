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
	GOOS=darwin $(GOBUILD) -o main.go -o $(RELEASEDIR)/macos/gstatic

build-linux:
	GOOS=linux $(GOBUILD) -o main.go -o $(RELEASEDIR)/linux/gstatic
	
build-win: setup
	GOOS=windows GOARCH=386 $(GOBUILD) -o main.go -o $(RELEASEDIR)/win/gstatic.exe
	
release: clean build-linux build-win build-macos