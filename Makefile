# Makefile for building Textstorm

# Application details
APP_NAME := Textstorm
APP_ID := se.stenstromen.textstorm
APP_VERSION := 1.0.0
ICON_FILE := big_icon.png

# Build targets
.PHONY: all amd64 arm64 universal packages clean

all: amd64 arm64 universal package

# Build for amd64
amd64:
	@CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o $(APP_NAME)_amd64

# Build for arm64
arm64:
	@GOOS=darwin GOARCH=arm64 go build -o $(APP_NAME)_arm64

# Create a universal binary
universal: amd64 arm64
	@lipo -create -output $(APP_NAME) $(APP_NAME)_amd64 $(APP_NAME)_arm64

# Package the application
package: universal
	@fyne package -os darwin -name "$(APP_NAME)" -icon $(ICON_FILE) --executable $(APP_NAME) --appID $(APP_ID) --appVersion $(APP_VERSION) --release

# Clean up build files
clean:
	@rm -f $(APP_NAME) $(APP_NAME)_amd64 $(APP_NAME)_arm64
