DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
PACKAGES = $(shell go list ./...)
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods \
				 -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

setup: 
	@go get github.com/axw/gocov/gocov
	@go get gopkg.in/matm/v1/gocov-html

all: deps format
	
cov:
	gocov test ./... | gocov-html > /tmp/coverage.html
	open /tmp/coverage.html

deps:
	@echo "--> Installing build dependencies"
	@go get -d -f -u -v ./...
	@echo $(DEPS) | xargs -n1 go get -d -f -u

test:
	@echo "--> Running go test"
	go list ./... | xargs -n1 go test

cover:
	@echo "--> Running go test --cover"
	go list ./... | xargs -n1 go test --cover

format:
	@echo "--> Running go fmt"
	@go fmt $(PACKAGES)

vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "--> Running go tool vet $(VETARGS) ."
	@go tool vet $(VETARGS) . ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for reviewal."; \
	fi

.PHONY: all cov deps test vet setup