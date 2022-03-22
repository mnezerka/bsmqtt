SOURCES := $(shell find . -name '*.go')

all: bsmqtt

bins: bsmqtt bsmqtt.exe bsmqtt.darwin

bsmqtt: $(SOURCES)
	GOOS=linux GOARCH=amd64 go build

bsmqtt.exe: $(SOURCES)
	GOOS=windows GOARCH=amd64 go build -o bsmqtt.exe

bsmqtt.darwin: $(SOURCES)
	GOOS=darwin GOARCH=amd64 go build -o bsmqtt.darwin
		
.PHONY: clean
clean:
	rm -f bsmqtt bsmqtt.exe bsmqtt.darwin