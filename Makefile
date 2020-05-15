all: build run 
build:
	go build -o bin/urlchecker *.go 

run:
	go run *.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o bin/urlchecker-linux-amd64 *.go
	GOOS=freebsd GOARCH=386 go build -o bin/urlchecker-freebsd-386 *.go

clean:
	rm -rf bin/urlchecker*
