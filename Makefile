TESTS?=$$(go list ./... | egrep -v "mock|docs")
BINARY=origin_backend
ENTRY=main.go

default: build

#host commands

install:
	go get

build: install
	go build -o $(BINARY) $(ENTRY)

run: build
	./$(BINARY)

clean:
	go clean
	rm -f $(BINARY) cover.out coverage.html

test:
	go test -v $(TESTS) -failfast -cover 

cover:
	go test -v $(TESTS) -failfast -coverprofile=cover.out
	go tool cover -html=cover.out -o coverage.html

docs: install
	swag init

#dockerized commands

docker/build:
	docker build -t origin_backend .

docker/build-test:
	docker build -t origin_backend_test . -f Dockerfile.test

docker/run: docker/build
	docker run --rm --name origin_backend -p 8080:8080 origin_backend

docker/clean:
	docker rmi -f origin_backend
	docker rmi -f origin_backend_test
	
docker/test: docker/build-test
	docker run --rm --name origin_backend_test origin_backend_test go test -v $(TESTS) -failfast