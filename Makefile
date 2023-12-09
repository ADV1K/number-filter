run: filter.go
	@go run . $(ARGS)

build: filter.go
	go build -o filter

test: filter.go filter_test.go
	go test .

clean:
	rm -f filter

.PHONY: run build test clean