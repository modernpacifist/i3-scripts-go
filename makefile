default:
	go build -o ./bin/ $(src)

build:
	find . -type f -name "*.go" | xargs -n 1 go build -o ./bin/

fmt:
	find . -type f -name "*.go" | xargs -n 1 go fmt

install:
	@if [ $$(id -u) != 0 ]; then echo "You must run install with root privileges"; exit 1; fi

	find . -type f -name "*.go" | xargs -n 1 go build -o ./bin/
	find ./bin/ -type f -executable -exec mv {} /bin/ \;
