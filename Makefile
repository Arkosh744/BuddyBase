regenerate_mocks:
	go generate -run="mockgen .*" -x ./...

test:
	go test -failfast -count=1 -v ./...

lint:
	go mod tidy
	golangci-lint run ./... --fix
