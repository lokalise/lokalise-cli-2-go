# defining variables
BINARY := lokalise2

run: build
	$(info running...)
	bin/$(BINARY)
.PHONY: run

release:
	goreleaser release
.PHONY: release

build: clean
	$(info building...)
	go build -o bin/$(BINARY) ./
.PHONY: build

clean:
	$(info cleaning...)
	rm -rf ./bin/*
	rm -rf ./dist/*
.PHONY: clean
