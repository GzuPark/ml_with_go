# Machine Learning with Go
- Go를 활용한 머신 러닝 책을 docker 환경으로 구성하여 실습
- Pachyderm을 사용하여 pipeline 구축 실습
- 특히, 버전이 맞지 않아 [책](http://acornpub.co.kr/book/ml-with-go)과 [코드](https://github.com/PacktPublishing/Machine-Learning-With-Go)를 그대로 실행하지 못했던 것을 수정

## Contents
<details>
<summary><strong>1장 : </strong>데이터 수집 및 구성</summary>

+ [Gopher 스타일로 데이터 처리하기](./ch01/handling_data_gopher_style/)
+ [CSV 파일](./ch01/csv_files/)
+ [JSON](./ch01/json/)
+ [SQL 유형 데이터베이스](./ch01/sql_like_databases/) (_Required PostgreSQL_)
+ [Caching](./ch01/caching/)

</details>

<details>
<summary><strong>2장 : </strong>행렬, 확률, 통계</summary>

+ [벡터](./ch02/vectors/)
+ [행렬](./ch02/matrices/)
+ [통계](./ch02/statistics/)
+ [확률 가설검정](./ch02/hypothesis/)

</details>

<details>
<summary><strong>3장 : </strong>평가 및 검증</summary>

+ [평가](./ch03/evaluation/)
+ [검증](./ch03/validation/)

</details>

<details>
<summary><strong>4장 : </strong>회귀분석</summary>

+ [선형 회귀분석](./ch04/linear_regression/)
+ [다중 선형 회귀분석](./ch04/multiple_regression/)
+ [비선형 회귀분석](./ch04/nonlinear_regression/)

</details>

<details>
<summary><strong>5장 : </strong>분류</summary>

+ [로지스틱 회귀분석](./ch05/logistic_regression/)
+ [k-최근접 이웃 모델](./ch05/knn/)
+ [의사결정 나무](./ch05/decision_tree/)
+ [나이브 베이즈](./ch05/naive_bayes/)

</details>

<details>
<summary><strong>6장 : </strong>클러스터링</summary>

+ [유사도 측정하기](./ch06/distance/)
+ [클러스터링 기법 평가하기](./ch06/evaluating/)
+ [k-평균 클러스터링](./ch06/kmeans/)

</details>

<details>
<summary><strong>7장 : </strong>시계열 분석 및 이상 감지</summary>

+ [시계열 데이터 표현하기](./ch07/representing_time_series/)
+ [시계열 통계](./ch07/ts_statistics/)
+ [자동 회귀 모델](./ch07/auto_regressive/)
+ [이상 감지](./ch07/anomaly_detection/)

</details>

<details>
<summary><strong>8장 : </strong>신경망 및 딥러닝</summary>

+ [간단한 신경망 구축하기](./ch08/simple_neural_network/)
+ [간단한 신경망 활용하기](./ch08/utilizing_simple_neural_network/)
+ [딥러닝](./ch08/deep_learning/) | [[How to use]](./ch08/deep_learning/README.md)

</details>

<details>
<summary><strong>9장 : </strong>분석 결과 및 모델 배포하기</summary>

+ [선형 회귀 모델 Dockerizing](./ch09/linear_regression/) | [[How to use]](./ch09/linear_regression/README.md)
+ [다중 회귀 모델 Dockerizing](./ch09/multiple_regression/) | [[How to use]](./ch09/multiple_regression/README.md)
+ [Docker로 모델 예측 수행하기](./ch09/predict_regression/) | [[How to use]](./ch09/predict_regression/README.md)
+ [머신 러닝 파이프라인 구축하기](./ch09/building_scalable_pipeline/) (_Required Pachyderm_) | [[How to use]](./ch09/building_scalable_pipeline/README.md)

</details>

## Environment
- `Docker Desktop >= 2.5.0.1` (_Recommended_)
- `Docker >= 19.03.8` (_Recommended_)
- `make` (_Required_)
- `wsl2` (_Only Windows_) :point_right: [설치 가이드](https://docs.microsoft.com/ko-kr/windows/wsl/install-win10)
    - Use __wsl__ terminal, not CMD or PowerShell on Windows
- `Golang >= 1.15` (_In the container_)
- `PostgreSQL >= 13.1` (_In the container_)
- `alpine >= 3.12` (_In the container_)
- `pachyderm ~= 1.11.7` (_only ch09_)

## How to use

#### Docker
- Compose up docker containers
    ```bash
    # Default
    (os)$ make up

    # With build go binaries
    (os)$ make up PKGS=1
    ```
- Start and enter the execution mode to the container
    ```bash
    (os)$ make run
    ```
- Stop docker containers
    ```bash
    (os)$ make stop
    ```
- Remove docker compose
    ```bash
    (os)$ make down
    ```

#### Build
```bash
(docker)$ cd tutorial_code_location
(docker)$ go build tutorial.go

# example
(docker)$ go build 01_read_csv_file.go
```

#### Execution
```bash
(docker)$ cd tutorial_code_location
(docker)$ ./tutorial

# example
(docker)$ ./01_read_csv_file
```

#### Aumotation
```bash
# build in anywhere
(docker)$ $MLGO/automation.sh build

# possible from ch01 to ch08
(docker)$ $MLGO/automation.sh build ch01

# clean up binary files in anywhere
(docker)$ $MLGO/automation.sh clean

# possible from ch01 to ch08 except deep_learning
(docker)$ $MLGO/automation.sh clean ch01

# execute binary files in anywhere
(docker)$ $MLGO/automation.sh run

# possible from ch01 to ch08 except deep_learning
(docker)$ $MLGO/automation.sh run ch01
```

## Reference
- [[Book] (번역) Go를 활용한 머신 러닝](http://acornpub.co.kr/book/ml-with-go)
- [[Book] Machine Learning With Go](https://www.packtpub.com/product/machine-learning-with-go/9781785882104)
- [[GitHub] Machine-Learning-With-Go](https://github.com/PacktPublishing/Machine-Learning-With-Go)
- [[Docs] Pachyderm](https://docs.pachyderm.com/latest/)
- [[Docs] Minikube](https://minikube.sigs.k8s.io/docs/)
