PACKAGE_NAME=github.com/parkr/silence-but-for-error

all: build test

build:
	go install $(PACKAGE_NAME)/...

test:
	go test $(PACKAGE_NAME)/...
