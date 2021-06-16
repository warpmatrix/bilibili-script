docker-image: main dockerfile
	docker build -t test .
	docker run -it test

main: main.go
	go build $^
