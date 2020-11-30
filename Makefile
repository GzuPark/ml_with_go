CONTAINER_NAME = mlgo
URL = https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz
PKGS = 0

up:
	@docker-compose up -d
	@docker container exec -it ${CONTAINER_NAME} \
		bash -c 'curl -L ${URL} | tar -C "/usr/local" -xz && ldconfig'
ifneq ($(PKGS), 0)
	@docker container exec -it ${CONTAINER_NAME} \
		bash -c './automation.sh build'
endif

run:
	@docker container start ${CONTAINER_NAME}
	@docker container exec -it ${CONTAINER_NAME} /bin/bash

stop:
	@docker-compose stop

down:
	@docker-compose down
