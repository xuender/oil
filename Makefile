all: lint

lint:
	# ,gosec,goerr113
	golangci-lint run --tests=false -E wsl,goerr113
