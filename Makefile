
all:
	make bulid
	make install

test:
	@go run ./src/ '12 Angry Men'
	@# prints:
	# > Search title
	# 25 titles found.
	# > Retrieve rating
	# First hit: '12 Angry Men', duration '1h36m'.
	# Rating: '9.0'.

.PHONY: build
build:
	rm -rf ./build/
	go build -o ./build/web_imdb ./src/main.go 

install:
	make build
	rm -f ~/.local/bin/web_imdb
	cp ./build/web_imdb ~/.local/bin/web_imdb

