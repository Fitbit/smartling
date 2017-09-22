FROM golang:1.9-alpine
ARG VERSION
ENV APP_DIR "$GOPATH/src/github.com/Fitbit/smartling"
WORKDIR $APP_DIR
COPY . $APP_DIR
RUN apk add --no-cache git && \
    go get -u -v github.com/Masterminds/glide && \
    glide install && \
    go build -ldflags "-X main.Version=${VERSION}" -o $GOPATH/bin/smartling ./cli/... && \
    apk del git && \
    rm -rf $GOPATH/src && \
    rm -rf $GOPATH/bin/glide
WORKDIR /usr/src/app
ENTRYPOINT ["/go/bin/smartling"]
