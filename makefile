SRCDIR = src
SRC = $(shell find $(SRCDIR) -name '*.go')

.PHONY: run clean

run: $(SRC) dockerfile docker-compose.yml
	docker-compose up -d --build

clean:
	docker-compose down
