TARGET = bin/main
BINDIR = bin
SRCDIR = src
SRC = $(shell find $(SRCDIR) -name '*.go')

docker-image: $(TARGET) dockerfile
	go test ./...
	docker build -t bilibili-script .
	docker run -it --env-file cookie.list -v $$(pwd)/config.yaml:/config.yaml bilibili-script

$(TARGET): $(SRC)
	go build -o $(BINDIR)/ ./...
	mv $(BINDIR)/$(SRCDIR) $(TARGET)

clean:
	@rm -rf bin
