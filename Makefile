all: clean install

clean:
	go clean -i github.com/kengibson1111/tour-of-go/...

install:
	go get -u golang.org/x/tour/pic
	go get -u golang.org/x/tour/reader
	go get -u golang.org/x/tour/tree
	go get -u golang.org/x/tour/wc
	go install ./...

