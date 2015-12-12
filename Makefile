build: bin/server

bin/server: src/server/server.go
	go install server

run: build
	bin/server

clean:
	rm bin/*

test:
	go test server

packages:
	go get -t server
