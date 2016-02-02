run:
	go run *.go

test:
	curl -X POST -d "{\"body\":\"test body\"}" http://localhost:9021/post

.PHONY: test
