VERSION=`cat VERSION`
all:
	gox -output="build/vinculum_{{.OS}}_{{.Arch}}" -os="linux"
release:
	git tag v${VERSION} 
	git push --tags
