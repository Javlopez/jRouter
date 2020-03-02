.PHONY: test

test:
	go test ./... -v


test-with-coverage:
	go test ./... -v -coverprofile coverage/cover.out

coverage: test-with-coverage
	go tool cover -html=coverage/cover.out

doc:
	godoc -http=:8081	