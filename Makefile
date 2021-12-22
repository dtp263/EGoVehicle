mock:
	rm -rf mocks
	./scripts/mock.sh

.PHONY: test
test:
	./scripts/test.sh
