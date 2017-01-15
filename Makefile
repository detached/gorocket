test:
	go test -v ./realtime
	go test -v ./rest

all:
	go build ./realtime
	go build ./rest
