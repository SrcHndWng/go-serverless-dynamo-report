.PHONY: build clean deploy test

build:
	cd report; go-assets-builder -p main -o assets.go assets/;
	cd report; env GOOS=linux go build -ldflags="-s -w" -o ../bin/report;

clean:
	rm -rf ./bin
	rm report/assets.go

deploy: clean build
	sls deploy --verbose

test:
	go test -v ./report/...
