# Machine Learning with Go
Go를 활용한 머신 러닝 책을 docker 환경으로 구성하여 실습

## Environment
- (_Recommend_) Docker >= 19.03.8

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

## Contents
<details>
<summary><strong>1장 : </strong> 데이터 수집 및 구성</summary>

+ [CSV 파일](./ch01/csv_files/)
+ [JSON](./ch01/json/)
+ [SQL 유형 데이터베이스](./ch01/sql_like_databases/) (_Required PostgreSQL_)

</details>

## Reference
- [[GitHub] Machine-Learning-With-Go](https://github.com/PacktPublishing/Machine-Learning-With-Go)
