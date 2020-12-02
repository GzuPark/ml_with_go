## 다중 회귀 모델 Dockerizing
`automation.sh`에서 작업이 진행되지 않도록 예외 처리를 한 상태이고, binary file로 build를 한 후 `goregtrain:multiple` docker image를 생성

#### How to use
```bash
# compile
(os)$ make compile
# build docker image
(os)$ make build
# train a model
(os)$ make train
# split diabetes dataset before train a model
(os)$ make train HAS_DATA=0
# remove a binary file
(os)$ make clean
# process all
(os)$ make
```
