CONTAINER_NAME = mlgo
WORKDIR = /go/src/mlgo/ch09/linear_regression
ALPINE_VERSION = 3.12.1
DIR :=

all: compile build train clean

compile:
	@docker container exec -w ${WORKDIR} ${CONTAINER_NAME} \
		bash -c 'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregtrain'

build:
	@docker build --force-rm=true --build-arg VERSION=${ALPINE_VERSION} -t goregtrain:single .

train:
ifeq ($(DIR),)
	$(error You must define the MODEL PATH)
endif
	@docker container run --rm \
		-v ${DIR}/model:/tmp/model \
		-v mlgo_data:/data \
		goregtrain:single \
		/goregtrain -inDir=/data -outDir=/tmp/model

clean:
	@docker container exec -w ${WORKDIR} ${CONTAINER_NAME} \
		bash -c 'rm goregtrain'