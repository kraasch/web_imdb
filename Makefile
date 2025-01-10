
all:
	make build
	make install

test:
	go test ./...

.PHONY: build
build:
	rm -rf ./build/
	go build -o ./build/web_imdb ./cmd/main.go 

install:
	rm -f ~/.local/bin/web_imdb
	cp ./build/web_imdb ~/.local/bin/web_imdb

