# defining variables
BINARY := lokalise2

run: build
	$(info running...)
	bin/$(BINARY)
.PHONY: run

release:
	git tag ${tag}
	git push origin ${tag}
	goreleaser release  --rm-dist --skip-validate --skip-publish
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

docs: build
	bin/lokalise2 gendocs
.PHONY: docs


push-alpine:
	docker build -f Dockerfile.alpine -t lokalise/lokalise-cli-2:alpine .
	docker push lokalise/lokalise-cli-2:alpine
.PHONY: push-alpine
