.PHONY: deps
deps:
	go mod download
	go mod tidy

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build:deps vet fmt test

.PHONY: test
test:
	go test -v ./...

.PHONY: update-deps
update-deps:
	go get buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go@latest
	go get github.com/plantoncloud/pulumi-module-golang-commons
	go get github.com/plantoncloud/kubernetes-crd-pulumi-types
