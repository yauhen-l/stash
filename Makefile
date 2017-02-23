all:
	go fmt
	go vet
	go test -v
	go install -v
