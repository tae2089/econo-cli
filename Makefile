BIN_DIR=./bin
BIN=econo-cli
BIN_WINDOWS=econo-cli.exe
BIN_DEBUG=$(BIN).debug
GCFLAGS_DEBUG="all=-N -l"
INSTALL_LOCATION=~/bin
WINDOWS_OS=windows
LINUX_OS=linux
MAC_OS=darwin
ARCH=amd64

.PHONY: build build-debug test install clean release bin-dir

build: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s/v.*/${{ github.event.release.tag_name }}/g" ./cmd/version.go; \
		go build -o $(BIN_DIR)/$(BIN); \
		git checkout -- ./cmd/version.go; \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-linux: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s/v.*/${{ github.event.release.tag_name }}/g" ./cmd/version.go; \
		GOOS=$(LINUX_OS) GOARCH=$(ARCH) go build -o $(BIN_DIR)/$(BIN) run/main.go; \
		tar -czvf $(BIN_DIR)/$(BIN).$(LINUX_OS)-$(ARCH).tar.gz $(BIN_DIR)/$(BIN); \
		git checkout -- ./cmd/version.go; \
		rm $(BIN_DIR)/$(BIN); \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-darwin: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s/v.*/${{ github.event.release.tag_name }}/g" ./cmd/version.go; \
		GOOS=$(MAC_OS) GOARCH=$(ARCH) go build -o $(BIN_DIR)/$(BIN) run/main.go; \
		tar -czvf $(BIN_DIR)/$(BIN).$(MAC_OS)-$(ARCH).tar.gz $(BIN_DIR)/$(BIN); \
		git checkout -- ./cmd/version.go; \
		rm $(BIN_DIR)/$(BIN); \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-windows: bin-dir
	if [ -z "$(shell git status --porcelain)" ]; then \
		sed -i "s/v.*/${{ github.event.release.tag_name }}/g" ./cmd/version.go; \
		GOOS=$(WINDOWS_OS) GOARCH=$(ARCH) go build -o $(BIN_DIR)/$(BIN_WINDOWS) run/main.go; \
		zip -9 -y $(BIN_DIR)/$(BIN).$(WINDOWS_OS)-$(ARCH).zip $(BIN_DIR)/$(BIN_WINDOWS); \
		git checkout -- ./cmd/version.go; \
		rm $(BIN_DIR)/$(BIN_WINDOWS); \
	else \
		echo Working directory not clean, commit changes; \
	fi

build-debug: bin-dir
	sed -i "s/v.*/${{ github.event.release.tag_name }}/g" ./cmd/version.go; \
	go build -o $(BIN_DIR)/$(BIN_DEBUG) -gcflags=$(GCFLAGS_DEBUG) run/main.go

bin-dir:
	mkdir -p $(BIN_DIR)

no-bin-dir:
	rm rf $(BIN_DIR)

test:
	go test -v ./...

install: build
	cp $(BIN_DIR)/$(BIN) $(INSTALL_LOCATION)/$(BIN)

release: build
	VERSION=$$($(BIN_DIR)/$(BIN) --version); \
	git tag $$VERSION;