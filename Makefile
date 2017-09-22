SHELL:=/bin/bash
ifdef TRAVIS_TAG
VERSION=$(TRAVIS_TAG)
else
VERSION=latest
endif
GOBUILD_ARGS:=-ldflags "-X main.Version=$(VERSION)"
BIN_NAME:=smartling
BUILD_DIR:=build
COVER_DIR:=coverage
DOCKER_TAG:=fitbit/${BIN_NAME}

.PHONY: clean build build-all pack-all fmt deps lint test bench cover cover-html coveralls docker

clean:
	@go clean ./...
	@- rm -rf ${COVER_DIR} ${BUILD_DIR}

build:
	@go build -o $$GOPATH/bin/$(BIN_NAME) ./cli/...

build-all:
	gox \
	$(GOBUILD_ARGS) \
	-os="linux darwin windows freebsd openbsd netbsd" \
	-arch="amd64 386 armv5 armv6 armv7 arm64" \
	-osarch="!darwin/arm64" \
	-output="${BUILD_DIR}/{{.OS}}-{{.Arch}}/${BIN_NAME}" ./cli/...

pack-all:
	@for dirname in $$(find ${BUILD_DIR} -mindepth 1 -maxdepth 1 -type d); do \
		basename=$$(basename $$dirname); \
		filename=${BIN_NAME}-${VERSION}-$$basename; \
		cp LICENSE $$dirname; \
		cp NOTICE $$dirname; \
		cp README.md $$dirname; \
		pushd $$dirname &> /dev/null; \
		tar -zcf ../$$filename.tar.gz ./; \
		zip -rq ../$$filename.zip ./; \
		popd &> /dev/null; \
	done

fmt:
	@go fmt ./...

deps:
	@go get -u -v github.com/axw/gocov/gocov
	@go get -u -v github.com/matm/gocov-html
	@go get -u -v github.com/wadey/gocovmerge
	@go get -u -v github.com/mattn/goveralls
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v github.com/mitchellh/gox
	@glide install

lint:
	@go vet ./...
	@golint $(go list ./...)

test:
	@go test -v ./...

bench:
	@go test ./... -bench . -benchtime 2s -benchmem

cover:
	@gocov test ./... | gocov report

cover-html:
	@- mkdir -p ${COVER_DIR}
	@gocov test $$(go list ./...) | gocov-html > ${COVER_DIR}/profile.html

coveralls:
	@- mkdir -p ${COVER_DIR}
	@for pkg in $$(go list ./...); do \
		go test $$pkg -coverprofile="${COVER_DIR}/$$(basename $$pkg)-profile.cov"; \
	done
	@gocovmerge ${COVER_DIR}/*-profile.cov > ${COVER_DIR}/profile.cov
	@goveralls -coverprofile=${COVER_DIR}/profile.cov -service=travis-ci

docker:
	docker build --force-rm -t ${DOCKER_TAG}:${VERSION} --build-arg VERSION=${VERSION} .
