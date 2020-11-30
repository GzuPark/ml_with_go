CONTAINER_NAME = mlgo
WORKDIR = /go/src/mlgo/ch09/predict_regression
ALPINE_VERSION = 3.12.1
MODEL = 0
DIR :=

all: compile build predict clean

compile:
	@docker container exec -w ${WORKDIR} ${CONTAINER_NAME} \
		bash -c 'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregpredict'

build:
	@docker build --force-rm=true --build-arg VERSION=${ALPINE_VERSION} -t goregpredict:latest .

predict:
ifeq ($(DIR),)
	$(error You must define the DIR variable as CURRENT PATH)
endif
ifeq ($(MODEL), 0)
	$(error You must select 'linear' or 'multiple')
endif
ifneq ($(MODEL), 0)
	@docker container run --rm \
		-v mlgo_data:/data \
		-v mlgo_model:/tmp/model \
		goregtrain:$(MODEL) \
		/goregtrain -inDir=/data -outDir=/tmp/model
endif
	@docker container run --rm \
		-v ${DIR}/attributes:/tmp/attributes \
		-v ${DIR}/result:/tmp/result \
		-v mlgo_model:/tmp/model \
		goregpredict:latest \
		/goregpredict -inModelDir=/tmp/model -inVarDir=/tmp/attributes -outDir=/tmp/result

clean:
	@docker container exec -w ${WORKDIR} ${CONTAINER_NAME} \
		bash -c 'rm goregpredict'
