all:
	go fmt
	go vet
	go clean
	GO15VENDOREXPERIMENT=0 godep go test 
	GO15VENDOREXPERIMENT=0 godep go build
	GO15VENDOREXPERIMENT=0 godep go install
