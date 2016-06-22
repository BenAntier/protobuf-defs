.PHONY: all proto test get-deps lint swagger

all:
	make -C logger all

lint:
	make -C logger lint

proto: lint
	make -C logger proto

swagger: proto
	make -C logger swagger

get-deps:
	make -C logger get-deps

test: proto swagger
	make -C logger test

check-proto-pre-compiled: proto
	@if git status --porcelain | grep -q -e '^.M'; then \
		echo ">>> ERROR: You must pre-compile the proto files, please run 'make proto', commit all changes and push"; \
		exit 1; \
	fi

ci: check-proto-pre-compiled test
