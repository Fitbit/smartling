PACKAGES:=$$(go list ./... | grep -v /vendor/)
ifdef TRAVIS_TAG
	VERSION=$(TRAVIS_TAG)
else
	VERSION=dev
endif
GOBUILD_ARGS:=-ldflags "-X main.Version=$(VERSION)"
BIN_NAME:=smartling
DIST_DIRS:=find * -type d -exec

.PHONY: clean build build-all fmt restore lint test bench cover cover-html coveralls readme

clean:
	@go clean $(PACKAGES)
	@- rm -rf dist
	@- rm -rf coverage
	@- rm -rf build

build:
	@go build -o $$GOPATH/bin/$(BIN_NAME) ./cli/...;

build-all:
	gox \
	$(GOBUILD_ARGS) \
	-os="linux darwin windows freebsd openbsd netbsd plan9" \
	-arch="amd64 386 armv5 armv6 armv7 arm64" \
	-osarch="!darwin/arm64" \
	-output="build/{{.OS}}-{{.Arch}}/${BIN_NAME}" ./cli/...

dist:
	cd build && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf ${BIN_NAME}-${VERSION}-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r ${BIN_NAME}-${VERSION}-{}.zip {} \; && \
	cd ..

fmt:
	@go fmt $(PACKAGES)

restore:
	@go get -u -v github.com/axw/gocov/gocov
	@go get -u -v github.com/matm/gocov-html
	@go get -u -v github.com/wadey/gocovmerge
	@go get -u -v github.com/mattn/goveralls
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/mitchellh/gox
	@glide install

lint:
	@for pkg in $(PACKAGES); do \
		go tool vet $$(basename $$pkg); \
		golint $$(basename $$pkg); \
	done

test:
	@go test -v $(PACKAGES)

bench:
	@go test $(PACKAGES) -bench . -benchtime 2s -benchmem

cover:
	@gocov test $(PACKAGES) | gocov report

cover-html:
	@- mkdir -p coverage
	@gocov test $(PACKAGES) | gocov-html > coverage/profile.html

coveralls:
	@- mkdir -p coverage
	@for pkg in $(PACKAGES); do \
		go test $$pkg -coverprofile="coverage/$$(basename $$pkg)-profile.cov"; \
	done
	@gocovmerge coverage/*-profile.cov > coverage/profile.cov
	@goveralls -coverprofile=coverage/profile.cov -service=travis-ci

readme:
	@npm run gitdown
