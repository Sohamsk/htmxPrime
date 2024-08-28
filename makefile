build:
	go build -o ./dist/main cmd/main.go

run: build
	./dist/main

clean:
	rm -rf ./dist