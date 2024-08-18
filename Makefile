CUR_DIR:=$(shell pwd)


build-test-env:
	docker build --tag network-test:latest .


privileged-tests: \
	build-test-env \
	test-routeAdd

test: build-test-env
	go test -v -failfast .

test-routeAdd:
	docker run -it --cap-add=NET_ADMIN -v $(CUR_DIR):/opt network-test:latest
