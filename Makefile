GOOSES= linux \
				darwin \
				windows

GOARCHS= 386 \
				 amd64 \
				 arm \
				 arm64

.PHONY: build
build:
	go build

.PHONY: install
install:
	go install

.PHONY: test
test:
	ginkgo -r -p -nodes 4 -progress

.PHONY: dist_linux
dist_linux:
	for arch in 386 amd64 arm arm64; do \
		CGO_ENABLED=0 GOOS=linux GOARCH=$${arch} go build -o dist/yf_linux_$${arch}; \
	done

.PHONY: dist_darwin
dist_darwin:
	for arch in 386 amd64; do \
		GOOS=darwin GOARCH=$${arch} go build -o dist/yf_darwin_$${arch}; \
	done

.PHONY: dist
dist: dist_linux dist_darwin

.PHONY: deps
deps:
	go mod vendor
