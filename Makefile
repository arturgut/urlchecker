all: build run 

# Local build
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

# Docker dev build and test
docker-build-dev:
	docker build -t urlchecker-dev -f Dockerfile.dev .

docker-run-dev:
	docker run urlchecker-dev:latest -d 

docker-run-test:
	docker run --rm --name urlchecker urlchecker-dev:latest go test

# Final build
docker-build:
	docker build -t mrsouliner/urlchecker . 

docker-push: 
	docker push mrsouliner/urlchecker:latest
