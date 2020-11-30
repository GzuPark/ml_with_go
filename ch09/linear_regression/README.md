## 모델 배포하기
`automation.sh`에서 작업이 진행되지 않도록 예외 처리를 한 상태이고, binary file로 build를 한 후 `goregtrain:single` docker image를 생성

#### How to use
```bash
# compile
(os)$ make compile
# build docker image
(os)$ make build
# train a model
(os)$ make train DIR=$(PWD)
# remove a binary file
(os)$ make clean
# process all
(os)$ make DIR=$(PWD)
```
