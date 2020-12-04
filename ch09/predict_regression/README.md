## Docker로 모델 예측 수행하기
`automation.sh`에서 작업이 진행되지 않도록 예외 처리를 한 상태이고, binary file로 build를 한 후 `goregpredict:latest` docker image를 생성

#### How to use
```bash
# compile
(os)$ make compile

# build docker image
(os)$ make build

# predict a model
(os)$ make predict DIR=${PWD}

# predict linear regression model
(os)$ make predict DIR=${PWD} MODEL=linear

# predict mutiple regression model
(os)$ make predict DIR=${PWD} MODEL=multiple

# remove a binary file
(os)$ make clean

# process all
(os)$ make DIR=${PWD}
```
