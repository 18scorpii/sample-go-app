# Sample Application In Golang

## Requirement

- Create a HR System to manage employees
	- CURD operation on employees
	- Store the data in the file system
	- REST API using JSON

## Execution Commands

- `go run main.go`
- `go clean -testcache && go test -v ./...`		 ( To Clear cache and run all test cases )
- `go test -v -run TestEmployeesUnmarshalling ./...`	( To Run a Specific Test Case )
- `go install` ( To create a binary for local setup under the GOBIN folder )
- `env GOOS=linux GOARCH=amd64 go build sample-go-app` ( To use this binary on a centos )

