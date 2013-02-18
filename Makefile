export GOPATH=$(CURDIR)/build-ext/:$(CURDIR)/build-my/

all:
	rm -fr build-my/
	mkdir -p build-ext/ build-my/src/logorezka
	ls -1 | grep -vE '^(build-ext|build-my|debian|Makefile)$$' | xargs cp -rt build-my/src/logorezka/
	go get -v -x -tags zmq_2_1 github.com/alecthomas/gozmq
	cd build-my/src/logorezka/ && go get -v && go build -v && go install -v
