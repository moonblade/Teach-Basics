build-counter:
	docker image rm counter || true
	docker build -t counter docker/resources/counter

bash-scripting-1:
	for x in {1..100}; do touch bash/resources/cleanup/simpleremove-$$x.png; done
	for x in {1..100}; do touch bash/resources/cleanup/file-$$x.csv; done
	for x in {1..100}; do touch bash/resources/cleanup/othercsvfile-$$x.csv; done
	for x in {1..100}; do touch bash/resources/cleanup/file-$$x.jpg; done
	for x in {1..100}; do touch bash/resources/cleanup/others-$$x.pdf; done
	for x in {1..100}; do touch bash/resources/cleanup/$$(openssl rand -base64 5)-$$x.pdf; done
	for x in {1..100}; do echo $$(openssl rand -base64 20) >> bash/resources/cleanup/$$(openssl rand -base64 5)-$$x.txt; done

clean:
	rm bash/resources/cleanup/*.*
