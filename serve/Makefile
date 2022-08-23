MAKEFLAGS += --no-builtin-rules

all: basedrez

basedrez: $(shell find **/*.go)
	go generate ./ent
	go build .

.PHONY: clean
clean:
	rm basedrez
