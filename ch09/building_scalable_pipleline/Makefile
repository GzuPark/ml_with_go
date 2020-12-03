CONTAINER_NAME = mlgo
WORKDIR = /go/src/mlgo/ch09/building_scalable_pipleline
GOOS :=
GOFILE :=
OS_LIST = linux darwin windows
MODEL = 0
CMD = $(shell echo GOOS=${GOOS} GOARCH=amd64 CGO_ENABLED=0 go build ${GOFILE})

compile:
ifeq ($(GOFILE),)
	$(error You must define the GOFILE)
endif

ifeq ($(filter $(GOOS),$(OS_LIST)),)
	$(error Yout must select specific OS in "${OS_LIST}")
endif

	@echo "Building... ${GOFILE}"
	@docker container exec -w ${WORKDIR} ${CONTAINER_NAME} \
		bash -c '${CMD}'
