.PHONY: build-mac build-linux build-windows

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o prondru-mac main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o prondru main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o prondru.exe main.go

all:
	GOOS=darwin GOARCH=amd64 go build -o prondru-mac main.go
	GOOS=linux GOARCH=amd64 go build -o prondru main.go
	GOOS=windows GOARCH=amd64 go build -o prondru.exe main.go
	mkdir bins
	mv prondru* bins

clean:
	go mod tidy
