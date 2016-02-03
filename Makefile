test:
	docker run --rm -it -v $$(pwd):/build/src quay.io/opsee/build-go:gogoopsee /bin/bash -c 'cd /build/src && ./build.sh'

.PHONY:
	test