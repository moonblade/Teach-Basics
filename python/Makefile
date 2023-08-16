env:
	docker run -it -v $$(pwd)/resources:/etc/resources ubuntu bash

py:
	docker run -it -v $$(pwd)/resources:/etc/resources ubuntu-python bash

image:
	docker build -t ubuntu-python resources

