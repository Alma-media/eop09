LINT_VERSION = 1.33.0
PROTODIR = proto

bootstrap:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v$(LINT_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go mod tidy
	@echo "Do not forget to install protoc C++ libraries manually"

lint:
	golangci-lint run --enable-all --disable lll,gochecknoglobals,dupl,interfacer,gochecknoinits,godox,funlen,gocognit,wsl,wrapcheck,exhaustivestruct,godot,testpackage,tparallel,paralleltest,gocritic

test:
	go vet ./...
	go test -v -race ./...

build-images:
	docker build -f Dockerfile.client -t alma/eop09-client:latest .
	docker build -f Dockerfile.server -t alma/eop09-server:latest .

proto-build: $(PROTODIR)/*.pb.go

%.pb.go: %.proto
	protoc --proto_path=. --go-grpc_out=paths=source_relative:. --go_out=paths=source_relative:. $^