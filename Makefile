PACKAGES:=$$(go list ./... | grep -v /vendor/)
ifdef TRAVIS_TAG
	VERSION=$(TRAVIS_TAG)
else
	VERSION=dev
endif
GOBUILD_ARGS:=-ldflags "-X main.Version=$(VERSION)"
BIN_NAME:=smartling

.PHONY: clean build build-all install fmt restore lint test bench cover cover-html coveralls readme

clean:
	@go clean $(PACKAGES)
	@- rm -rf dist
	@- rm -rf coverage
	@- rm -rf build

build:
	@go build -o $$GOPATH/bin/$(BIN_NAME) ./cli/...;

install:
	@go install $(PACKAGES)

build-all:
	@for GOOS in darwin linux; do \
		for GOARCH in 386 amd64; do \
			echo "Building $$GOOS-$$GOARCH..."; \
			FULL_BIN_NAME=$(BIN_NAME)-$$GOOS-$$GOARCH; \
			GOOS=$$GOOS GOARCH=$$GOARCH go build $(GOBUILD_ARGS) -o build/$$FULL_BIN_NAME ./cli/...; \
			cd build; \
			cp $$FULL_BIN_NAME $(BIN_NAME); \
			tar -zcf "$$FULL_BIN_NAME.tar.gz" $(BIN_NAME); \
			rm -rf $(BIN_NAME); \
			cd ..; \
		done \
	done

fmt:
	@go fmt $(PACKAGES)

restore:
	@go get -u -v github.com/axw/gocov/gocov
	@go get -u -v github.com/matm/gocov-html
	@go get -u -v github.com/wadey/gocovmerge
	@go get -u -v github.com/mattn/goveralls
	@go get -u -v github.com/golang/lint/golint
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
