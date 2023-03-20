build-counter:
	docker image rm counter || true
	docker build -t counter docker/resources/counter
