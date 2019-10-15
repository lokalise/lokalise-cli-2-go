# defining variables
GO ?= go
BINARY := lokalise2
GIT_REV := $$(git rev-parse HEAD)
PLATFORMS=darwin linux windows freebsd
ARCHITECTURES=386 amd64

run: build
	$(info running...)
	bin/$(BINARY)
.PHONY: run

build_all:
	$(foreach GOOS, $(PLATFORMS), \
        $(foreach GOARCH, $(ARCHITECTURES), \
            $(shell \
            export GOOS=$(GOOS); \
            export GOARCH=$(GOARCH); \
            EXT=""; \
            if [ ${GOOS} == "windows" ]; then EXT=".exe"; fi; \
            go build -v -o bin/$(BINARY)$$EXT; \
            tar -C bin/ -cf dist/lokalise2-${VERSION}-${GOOS}-${GOARCH}.tgz $(BINARY)$$EXT)))

build: clean
	$(info building...)
	go build -o bin/$(BINARY) ./
.PHONY: build

clean:
	$(info cleaning...)
	rm -rf ./bin/*
.PHONY: clean
