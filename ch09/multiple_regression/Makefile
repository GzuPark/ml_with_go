CONTAINER_NAME = mlgo
ALPINE_VERSION = 3.12.1
WORKDIR = /go/src/mlgo/ch09/multiple_regression
DATASPLITDIR = /go/src/mlgo/ch03/validation
HAS_DATA = 1

all: compile build train clean

compile:
	@docker container exec -w ${WORKDIR} ${CONTAINER_NAME} \
		bash -c 'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregtrain'

build:
	@docker build --force-rm=true --build-arg VERSION=${ALPINE_VERSION} -t goregtrain:multiple .

train:
ifeq ($(HAS_DATA), 0)
	@docker container exec -w ${DATASPLITDIR} ${CONTAINER_NAME} \
		bash -c 'go build 01_split_dataset.go'
	@docker container exec -w ${DATASPLITDIR} ${CONTAINER_NAME} \
		bash -c './01_split_dataset'
endif
	@docker container run --rm \
		-v mlgo_storage:/storage \
		goregtrain:multiple \
		/goregtrain -inDir=/storage/data -outDir=/storage/model

clean:
	@docker container exec -w ${WORKDIR} ${CONTAINER_NAME} \
		bash -c 'rm goregtrain'
