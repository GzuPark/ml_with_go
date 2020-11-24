# Machine Learning with Go
Go를 활용한 머신 러닝 책을 docker 환경으로 구성하여 실습

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

## Environment
- `Docker >= 19.03.8` (_Recommended_)
- `Golang >= 1.15`

#### Windows
- Required
    - `wsl2`
    - `chocolatey`
    - `make`

#### MacOS
- Not to need to prepare

## How to use

#### Docker
- Create docker container
    ```bash
    # Windows
    (os)$ make run ACCOUNT=[ GitHub ] DIR=${PWD} [ GO_VERSION CONTAINER_NAME ]
    # MacOS
    (os)$ make run ACCOUNT=[ GitHub ] DIR=$(pwd) [ GO_VERSION CONTAINER_NAME ]
    ```
- Start and attach to the container
    ```bash
    (os)$ make start [ CONTAINER_NAME ]
    ```
- Stop the container
    ```bash
    (os)$ make stop [ CONTAINER_NAME ]
    ```
- Remove the container
    ```bash
    (os)$ make del [ CONTAINER_NAME ]
    ```

#### Installation packages
```bash
(docker)$ go get package_url
# example
(docker)$ go get github.com/go-gota/gota/...
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

## Reference
- [[GitHub] Machine-Learning-With-Go](https://github.com/PacktPublishing/Machine-Learning-With-Go)
