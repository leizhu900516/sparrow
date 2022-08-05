GO_CMD=go
SPARROW_BIN=sparrow_bin
SPARROW_BIN_PATH=./output/bin
SPARROW_ETC_PATH=./output/config

.PHONY: all
all: clean build install

.PHONY: linux
linux: clean build-linux install

.PHONY: build
build:
	@echo "build sparrow start >>>"
	GOPROXY=https://goproxy.io $(GO_CMD) mod tidy
	$(GO_CMD) build -o $(SPARROW_BIN) ./main.go
	@echo ">>> build sparrow complete"

.PHONY: install
install:
	@echo "install sparrow start >>>"
	mkdir -p $(SPARROW_BIN_PATH)
	mv $(SPARROW_BIN) $(SPARROW_BIN_PATH)/sparrow
	mkdir -p $(SPARROW_ETC_PATH)
	cp ./config/config.ini $(SPARROW_ETC_PATH)/config.ini
	@echo ">>> install sparrow complete"

.PHONY: clean
clean:
	@echo "clean start >>>"
	rm -fr ./output
	rm -f $(SPARROW_BIN)
	@echo ">>> clean complete"

.PHONY: build-linux
build-linux:
	@echo "build-linux start >>>"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_CMD) build -o $(SPARROW_BIN) ./main.go
	@echo ">>> build-linux complete"
