all:
	gox -output="build/vinculum_{{.OS}}_{{.Arch}}" -os="linux"
