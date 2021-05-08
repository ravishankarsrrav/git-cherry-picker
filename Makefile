.PHONY: help run build install

# target: help - Display callable targets.
help:
	@egrep "^# target:" [Mm]akefile


# target: run - start the gcpk locally
run:
	go run cmd/main.go


# target: install - install (and update) go packages
install:
	go install

# target: build - builds the gcpk tool
build:
	cd cmd && go build -o ../bin/gcpk


