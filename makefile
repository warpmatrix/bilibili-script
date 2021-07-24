SRCDIR = src
SRC = $(shell find $(SRCDIR) -name '*.go')
BINDIR = bin
BIN = app

.PHONY: run clean

run: $(SRC) dockerfile docker-compose.yml
	docker-compose up -d --build

$(BINDIR)/$(BIN): $(SRC)
	go build -o $(BINDIR)/ ./...
	mv $(BINDIR)/$(SRCDIR) $(BINDIR)/$(BIN)

test: $(SRC)
	go test -v ./...

clean:
	rm -rf $(BINDIR)
	docker-compose down
