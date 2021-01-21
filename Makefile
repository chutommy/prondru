.PHONY: build-mac build-linux build-windows all clean

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o prondru-mac main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o prondru main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o prondru.exe main.go

all:
	go mod tidy
	GOOS=darwin GOARCH=amd64 go build -o prondru-mac main.go
	GOOS=linux GOARCH=amd64 go build -o prondru main.go
	GOOS=windows GOARCH=amd64 go build -o prondru.exe main.go
	rm -rf bins
	mkdir bins
	mv prondru* bins

clean:
	go mod tidy
