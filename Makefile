.PHONY: build run clean

build:
	go build -o service ./service

run:
	go run ./service

clean:
	rm -f service





