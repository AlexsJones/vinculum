VERSION=`cat VERSION`
all:
	gox -output="build/vinculum_{{.OS}}_{{.Arch}}" -os="linux"
release:
	git tag v${VERSION} 
	git push --tags
generate:
	protoc --proto_path=pkg/proto pkg/proto/*.proto --go_out=plugins=grpc:.
docker-build:
	docker build -f builder.dockerfile .