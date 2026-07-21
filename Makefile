
TARGET = muxt
TARGETDIR = build
SRCS = $(shell fd -e .go)

PREFIX ?= /usr/local
BINDIR = $(PREFIX)/bin

.PHONY: install clean

all: $(TARGETDIR)/$(TARGET)

$(TARGETDIR)/$(TARGET): $(SRCS)
	mkdir -p $(TARGETDIR)
	go build -o $@ cmd/main.go

install:
	cp $(TARGETDIR)/$(TARGET) $(BINDIR)/

uninstall:
	rm $(BINDIR)/$(TARGET) 

clean:
	rm -r $(TARGETDIR)
