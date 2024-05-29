run: 
	go run .

builds: 
	go build -o ./build/main

test: 
	go test

clean:
	rm main