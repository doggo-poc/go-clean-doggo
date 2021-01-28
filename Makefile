build:
	go build DoggosPkg 
test:
	go test ./... -coverprofile cover.out
calculate-coverage:
	go tool cover -func cover.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'
