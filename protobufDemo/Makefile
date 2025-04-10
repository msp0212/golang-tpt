# make sure curr directory is in $GOPATH
compile:
	protoc -I=. --go_out=. appAccessRecord.proto
	go build .

run: compile
	./protobufDemo
