build:
	go build -v ./api
	go build -v ./realtime
	go build -v ./rest

test:
	go build ./common_testing
	go test -v ./realtime
	go test -v ./rest

get:
	go get -v -t ./common_testing
	go get -v -t ./realtime
	go get -v -t ./rest

coverage:
	go build ./common_testing
	go test -v -covermode=count -coverprofile=realtime.cov ./realtime
	go test -v -covermode=count -coverprofile=rest.cov ./rest
	sed -i '/mode: count/d' realtime.cov && cat rest.cov realtime.cov > all.cov
	goveralls -coverprofile=all.cov -service=travis-ci
	
