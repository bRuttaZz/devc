INSTALL_PATH = /usr/local/bin
ARCH=amd64

VERSION:=$(shell cat VERSION | tr -d ' ') 

compile: devc.go cmd internal go.mod 
	- @echo "[devc] compiling for : $(VERSION)."
	- GOOS=linux GOARCH=$(ARCH) go build -v -o "devc-$(strip $(VERSION))-linux-$(strip $(ARCH))" -ldflags="-s -w -X main.version=$(VERSION)"

install: devc
	- sudo cp devc $(INSTALL_PATH)/devc
