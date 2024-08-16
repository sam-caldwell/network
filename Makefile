CUR_DIR:=$(shell pwd)


build-test-env:
	docker build --tag network-test:latest .

privileged-tests: \
	build-test-env \
	test-routeAdd

test: privileged-tests
	@echo 'ok'

test-routeAdd:
	docker run -it --cap-add=NET_ADMIN -v $(CUR_DIR):/opt network-test:latest go run examples/addRoute/main.go

# ToDo: add more tests here