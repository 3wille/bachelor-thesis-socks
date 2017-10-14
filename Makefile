build:
	go build -o main main.go
	-ssh ovh 'killall main'
	scp main ovh:~/
