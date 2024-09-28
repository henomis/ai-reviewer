.PHONY: build clean

build:
	go build -v -o build/ai-reviewer .

clean:
	rm -fr build
