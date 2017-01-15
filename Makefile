all:
	go build ./realtime
	go build ./rest

test:
	go test -v ./realtime
	go test -v ./rest

coverage:
	go test -covermode=count -coverprofile=realtime.cov ./realtime
	go test -covermode=count -coverprofile=rest.cov ./rest
	sed -i '/mode: count/d' realtime.cov && cat rest.cov realtime.cov > all.cov
	goveralls -coverprofile=all.cov -service=travis-ci
	
