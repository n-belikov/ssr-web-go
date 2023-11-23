CWD=$(shell pwd)
GOPATH := $(CWD)

PROJECT_NAME=ssr-web

linux:
	GOOS=linux go build -o build/$(PROJECT_NAME)-linux main.go
mac:
	GOOS=darwin go build -o build/$(PROJECT_NAME)-mac main.go
windows:
	GOOS=windows go build -o build/$(PROJECT_NAME)-windows.exe main.go

all: linux mac windows
