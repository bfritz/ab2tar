VERSION = 0.0.1-SNAPSHOT
LDFLAGS = -s -w -X main.version=${VERSION}

sources := ab2tar.go

all: compile

deps:
	go get github.com/alexflint/go-arg

compile: $(sources) deps
	GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS}"
