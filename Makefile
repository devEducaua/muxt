
TARGET = muxt
TARGETDIR = build
SRCS = $(shell fd -e .go)

.PHONY: clean

all: $(TARGETDIR)/$(TARGET)

$(TARGETDIR)/$(TARGET): $(SRCS)
	mkdir -p $(TARGETDIR)
	go build -o $@ cmd/main.go

clean:
	rm -r $(TARGETDIR)
