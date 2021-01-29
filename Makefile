build:
	go build DoggosPkg 
ci-coverage:
	go test -v  -coverpkg=./... ./... -coverprofile cover.out 
	go tool cover -func cover.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'
