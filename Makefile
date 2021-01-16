VERSION=$(shell cat version)
REVISION=$(shell git rev-parse --short HEAD)

define linker_flags
-X github.com/mwlng/aws-go-clients/versions.Version=$(VERSION)\
-X github.com/mwlng/aws-go-clients/versions.Revision=$(REVISION) \
-extldflags "-static"
endef

all: linter

linter:
	golangci-lint run -p style -p performance -p format -p complexity -p bugs -p unused -D dupl -D gochecknoglobals -D gochecknoinits -D lll -D forbidigo -D exhaustivestruct -D gomnd ./...

#test:
#	go test -v -race $(shell go list ./... | grep -v tests/endpoint)

