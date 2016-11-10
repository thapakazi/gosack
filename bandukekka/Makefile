all:
	go run main.go
arm:
	env GOOS=linux GOARCH=arm go build && scp githubapi rpi:~/.bin
compile:
	go build 
