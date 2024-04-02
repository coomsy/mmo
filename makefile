.PHONY: serve tidy test

serve:
	MMO_DEBUG=true go run cmd/api/api.go
tidy:
	go mod tidy && go mod vendor

