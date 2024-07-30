all: lorenzod-env

lorenzod-env: lorenzod-rmi
	docker build --tag lorenzo/lorenzod -f lorenzod-env/Dockerfile \
		$(shell git rev-parse --show-toplevel)

lorenzod-dlv: lorenzod-rmi
	docker build --tag lorenzo/lorenzod -f lorenzod-dlv/Dockerfile \
		$(shell git rev-parse --show-toplevel)

lorenzod-rmi:
	docker rmi lorenzo/lorenzod 2>/dev/null; true

.PHONY: all lorenzod-env lorenzod-dlv lorenzod-rmi
