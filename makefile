docker-image: main dockerfile
	docker build -t test .
	docker run -it --env-file env.list test

main: main.go
	go build $^
