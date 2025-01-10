
all:
	make build
	make install

test:
	@go run ./cmd/ '12 Angry Men'
	@# prints:
	# > Search title
	# 25 titles found.
	# > Retrieve rating
	# First hit: '12 Angry Men' (1957), duration '1h36m'.
	# Rating: '9.0'.

.PHONY: build
build:
	rm -rf ./build/
	go build -o ./build/web_imdb ./cmd/main.go 

install:
	rm -f ~/.local/bin/web_imdb
	cp ./build/web_imdb ~/.local/bin/web_imdb

