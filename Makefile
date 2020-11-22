GO_VERSION = 1.15.5
CONTAINER_NAME = mlgo
DIR := 
	
win:
ifeq ($(DIR),)
	$(error You must define the CURRENT PATH)
endif
		@docker container run -it --name=${CONTAINER_NAME} -v ${DIR}:/go/src golang:${GO_VERSION} /bin/bash
start:
	@docker start ${CONTAINER_NAME}
	@docker attach ${CONTAINER_NAME}
stop:
	@docker stop ${CONTAINER_NAME}
del:
	@docker rm -f ${CONTAINER_NAME}
