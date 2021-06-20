TARGET = bin/main
SRCDIR = src
SRC = $(shell find $(SRCDIR) -name '*.go')

docker-image: $(TARGET) dockerfile
	docker build -t bilibili-script .
	docker run -it --env-file env.list bilibili-script

$(TARGET): $(SRC)
	go build -o $(TARGET) ./$(SRCDIR)

clean:
	@rm -rf bin
