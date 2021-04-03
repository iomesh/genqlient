example:
	go generate ./...
	go run ./example/cmd/example/main.go

check:
	go test ./...

genqlient.png: genqlient.svg
	convert -density 600 -background transparent "$<" "$@"

.PHONY: example
