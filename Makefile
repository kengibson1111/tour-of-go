all: clean install

clean:
	go clean -i github.com/kengibson1111/tour-of-go/...

install:
	go get -u github.com/kengibson1111/tour-of-go/...

