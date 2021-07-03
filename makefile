SRCDIR = src
SRC = $(shell find $(SRCDIR) -name '*.go')
BINDIR = bin
APP = $(BINDIR)/main
IMAGE = bilibili-script

.PHONY: run clean

run: $(APP) dockerfile docker-compose.yml
	docker-compose up -d --build

$(APP): $(SRC)
	go test ./...
	go build -o $(BINDIR)/ ./...
	mv $(BINDIR)/$(SRCDIR) $(APP)

clean:
	@rm -rf $(BINDIR)
	docker-compose down
