PROTO_DIR = proto
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
BIN = proto-go-course

## Adds Git hooks to correct location and configures Git to use SSH over HTTPS
config:
	git config --global url."git@github.com:".insteadOf "https://github.com/"

.PHONY: mod
mod: config ## Initialise the go module
	rm -rf go.sum go.mod vendor
	git config --global --replace-all url."git@github.com:".insteadOf "https://github.com/"
	go mod init github.com/samme4life/proto-go-course
	go mod tidy

install: ## installs dependencies
	go mod download
	go mod vendor

# Fetch project & Go dependencies
# This target will only work if the user has go1.18.5 on their PATH
# How to configure specific go version is explained here: https://go.dev/doc/manage-install
.PHONY: deps
deps: config
	go mod tidy
	go mod download
	go mod vendor

	echo deps done

build: 	generate
	go build .

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto

#bump: generate
#	go get -u ./...

clean:
	rm ${PROTO_DIR}/*.pb.go
	rm ${BIN}