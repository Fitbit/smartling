SHELL:=/bin/bash
PACKAGES:=$$(go list ./... | grep -v /vendor/)
ifdef TRAVIS_TAG
	VERSION=$(TRAVIS_TAG)
else
	VERSION=dev
endif
GOBUILD_ARGS:=-ldflags "-X main.Version=$(VERSION)"
BIN_NAME:=smartling
BUILD_DIR:=build
COVER_DIR:=coverage

.PHONY: clean build build-all dist fmt restore lint test bench cover cover-html coveralls readme

clean:
	@go clean $(PACKAGES)
	@- rm -rf ${COVER_DIR}
	@- rm -rf ${BUILD_DIR}

build:
	@go build -o $$GOPATH/bin/$(BIN_NAME) ./cli/...;

build-all:
	gox \
	$(GOBUILD_ARGS) \
	-os="linux darwin windows freebsd openbsd netbsd" \
	-arch="amd64 386 armv5 armv6 armv7 arm64" \
	-osarch="!darwin/arm64" \
	-output="${BUILD_DIR}/{{.OS}}-{{.Arch}}/${BIN_NAME}" ./cli/...

dist:
	@for dirname in $$(find ${BUILD_DIR} -mindepth 1 -maxdepth 1 -type d); do \
		basename=$$(basename $$dirname); \
		filename=${BIN_NAME}-${VERSION}-$$basename; \
		cp LICENSE $$dirname; \
		cp README.md $$dirname; \
		pushd $$dirname &> /dev/null; \
		tar -zcf ../$$filename.tar.gz ./; \
		zip -rq ../$$filename.zip ./; \
		popd &> /dev/null; \
	done

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
	@gocov test $(PACKAGES) | gocov-html > ${COVER_DIR}/profile.html

coveralls:
	@- mkdir -p coverage
	@for pkg in $(PACKAGES); do \
		go test $$pkg -coverprofile="${COVER_DIR}/$$(basename $$pkg)-profile.cov"; \
	done
	@gocovmerge ${COVER_DIR}/*-profile.cov > ${COVER_DIR}/profile.cov
	@goveralls -coverprofile=${COVER_DIR}/profile.cov -service=travis-ci

readme:
	@npm run gitdown
