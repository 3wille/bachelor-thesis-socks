build_and_copy: build
	-ssh ovh 'killall main'
	scp main ovh:~/

build:
	go build -o main main.go
