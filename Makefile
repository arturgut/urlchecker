all: build run 
build:
	go build -o bin/urlchecker *.go
	chmod +x bin/urlchecker

run:
	go run *.go

compile:
	echo "Compiling for linux-amd64"
	GOOS=linux GOARCH=amd64 go build -o bin/urlchecker-linux-amd64 *.go

clean:
	rm -rf bin/urlchecker*

test:
	go test

docker-build-dev:
	docker build -t urlchecker-dev -f Dockerfile.dev .

docker-run-dev:
	docker exec -it -u root jenkins /bin/bash

docker-build:
	docker build -t mrsouliner/urlchecker . 

docker-run:
	docker run -p 8091:8091 mrsouliner/urlchecker:latest

docker-push: 
	docker push mrsouliner/urlchecker:latest
