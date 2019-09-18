# defining variables
GO ?= go
OUT := bin/lokalise
GIT_REV := $$(git rev-parse HEAD)

run: build
	$(info running...)
	$(OUT)
.PHONY: run

# extended version:
# go build -ldflags "-X ./config.Version=$(GIT_REV)" -o $(OUT) ./
build: clean
	$(info building...)
	go build -o $(OUT) ./
.PHONY: build

clean:
	$(info cleaning...)
	rm -rf ./bin/*
.PHONY: clean

# gen.swagger
# gen.godoc