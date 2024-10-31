export GOOS
export GOARCH

all: build-linux-amd64 build-linux-arm64 build-windows-amd64

build-linux-amd64:
	@echo Building for Linux AMD64...
	$(MAKE) build GOOS=linux GOARCH=amd64

build-linux-arm64:
	@echo Building for Linux ARM64...
	$(MAKE) build GOOS=linux GOARCH=arm64

build-windows-amd64:
	@echo Building for Windows AMD64...
	$(MAKE) build GOOS=windows GOARCH=amd64

build:
	@echo Building with GOOS=$(GOOS) GOARCH=$(GOARCH)
ifeq ($(GOOS),windows)
	go build -o bin/$(GOOS)_$(GOARCH)/csphash.exe ./cmd
else
	go build -o bin/$(GOOS)_$(GOARCH)/csphash ./cmd
endif

clean:
	@echo Cleaning build artifacts...
ifeq ($(OS),Windows_NT)
	if exist bin rmdir /S /Q bin
else
	rm -rf bin
endif

compress:
	@echo Compressing build artifacts...
ifeq ($(OS),Windows_NT)
	@for /D %%d in (bin\*) do \
		PowerShell -Command "Compress-Archive -Path '%%d\*' -DestinationPath '%%d.zip'"
else
	@for dir in bin/*; do \
		if [ -d "$$dir" ]; then \
			zip -r "$$dir.zip" "$$dir"; \
		fi; \
	done
endif

.PHONY: all build-linux build-windows build clean