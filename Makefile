.PHONY: all proto test get-deps lint

all:
	make -C logger all

lint:
	make -C logger lint

proto: lint
	make -C logger proto

get-deps:
	make -C logger get-deps

test: proto
	make -C logger test

check-proto-pre-compiled: proto
	@if git status --porcelain | grep -q -e '^.M'; then \
		echo ">>> ERROR: You must pre-compile the proto files, please run 'make proto', commit all changes and push"; \
		exit 1; \
	fi
