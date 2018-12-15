MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "1.0")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
LDFLAGS:=-ldflags "-X github.com/alamin-mahamud/gapi/cmd/api.Version=${VERSION}"

default: run

test:
	go test -p=1 -cover -covermode=count ${PACKAGES}

cover:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
		echo ${pkg}; \
		go test -coverprofile=coverage.out -covermode=count ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

run:
	go run ${LDFLAGS} main.go

build: clean
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -a -o gapi main.go

clean:
	rm -f gapi coverage.out coverage-all.out