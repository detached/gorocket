build:
	go build ./api
	go build ./realtime
	go build ./rest

test:
	go build ./common_testing
	go test -v ./realtime
	go test -v ./rest

get:
	go get -v ./common_testing
	go get -v ./realtime
	go get -v ./rest

coverage:
	go build ./common_testing
	go test -covermode=count -coverprofile=realtime.cov ./realtime
	go test -covermode=count -coverprofile=rest.cov ./rest
	sed -i '/mode: count/d' realtime.cov && cat rest.cov realtime.cov > all.cov
	goveralls -coverprofile=all.cov -service=travis-ci
	
