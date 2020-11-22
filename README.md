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
- TBC

## How to use


#### Docker
- Create docker container
    ```bash
    (os)$ account=github_account
    (os)$ make run ACCOUNT=${account} DIR=${PWD} [ GO_VERSION CONTAINER_NAME ]
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
(docker)$ go get -u gonum.org/v1/gonum/...
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
