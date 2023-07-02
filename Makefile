PROTO_DIR = ./internal/proto
BIN = proto-go-course

## Adds Git hooks to correct location and configures Git to use SSH over HTTPS
config:
	git config --global url."git@github.com:".insteadOf "https://github.com/"

.PHONY: mod
mod: config ## Initialise the go module
	rm -rf go.sum go.mod vendor
	git config --global --replace-all url."git@github.com:".insteadOf "https://github.com/"
	go mod init proto-go-course
	go mod tidy

install: ## installs dependencies
	go mod download
	go mod vendor

# Fetch project & Go dependencies
.PHONY: deps
deps: config
	go mod tidy
	go mod download
	go mod vendor

	echo deps done

build: 	generate
	go build .

generate:
	protoc --go_out=paths=source_relative:. ${PROTO_DIR}/*.proto

clean:
	rm -f ${PROTO_DIR}/*.pb.go
	rm -f ${BIN}
	rm -rf ./vendor

.PHONY: create_exe
create_exe: clean generate mod install
	go build .


run_exe:
	./run.sh
