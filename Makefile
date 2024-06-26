test:
	@go fmt ./...
	@go test -cover ./...
	
coverage:
	@go fmt ./...
	@go test -coverprofile cover.out ./... 
	@go tool cover -html=cover.out
	@rm cover.out