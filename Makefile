
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

