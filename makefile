TARGET = bin/main
SRCDIR = src
SRC = $(shell find $(SRCDIR) -name '*.go')

docker-image: $(TARGET) dockerfile
	docker build -t bilibili-script .
	docker run -it --env-file cookie.list -v $$(pwd)/config.yaml:/config.yaml bilibili-script

$(TARGET): $(SRC)
	go build -o $(TARGET) ./$(SRCDIR)

clean:
	@rm -rf bin
