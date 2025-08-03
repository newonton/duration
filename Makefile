APP_NAME=duration
GOPATH=$(shell go env GOPATH)
INSTALL_DIR=$(GOPATH)/bin

.PHONY: all build run clean install uninstall

all: build

build:
	go build -o $(APP_NAME) main.go

install: build
	install -m 755 $(APP_NAME) $(INSTALL_DIR)/$(APP_NAME)
	@rm -f $(APP_NAME)
	@echo "Installed to $(INSTALL_DIR)/$(APP_NAME)"

uninstall:
	@rm -f $(INSTALL_DIR)/$(APP_NAME)
	@echo "Removed $(INSTALL_DIR)/$(APP_NAME)"

clean:
	@rm -f $(APP_NAME)
	@echo "Cleaned build"
