VERSION = 0.0.1-SNAPSHOT
LDFLAGS = -s -w -X main.version=${VERSION}

sources := ab2tar.go

all: compile

deps:

compile: $(sources) deps
	GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS}"
