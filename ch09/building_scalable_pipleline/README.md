## 확장 가능하고 재현 가능한 머신 러닝 파이프라인 구축하기
[Pachderm](https://www.pachyderm.com/)은 [Kubernetes](https://kubernetes.io/)를 사용하는 빅데이터 플랫폼으로 데이터 이력 관리가 용이하고, Docker image를 사용한 파이프라인을 구성하는 것이 특징이다. 로컬 환경에서 실습하기 위해 Kubernetes를 대신하여 [Minikube](https://minikube.sigs.k8s.io/docs/)를 이용하고, 앞서 만들었던 [선형 회귀 모델](../linear_regression/README.md), [다중 회귀 모델](../multiple_regression/README.md) 그리고 [회귀 모델 예측](../predict_regression/README.md)에서 만든 이미지들을 사용한다.

###### _Updated: 2020/12/02_

### Contents
- [Environment](#environment)
    - [All](#all)
    - [MacOS](#macos)
    - [Windows or Debian-based Linux](#windows-or-debian-based-linux)
- [How to use](#how-to-use)
    - [Deploy](#deploy)
    - [Run with Go files](#run-with-go-files)
    - [Run with pachctl](./pachctl/README.md)
    - [Useful pachctl command](#useful-pachctl-command)
    - [Undeploy](#undeploy)

### Environment
- `Docker Desktop >= 2.5.0.1` (_Recommended_)
- `minikube ~= 1.15.1` (_Recommended_)
- `pachyderm ~= 1.11.7` (_Required_)
    - Update가 활발하게 진행되므로 version에 따라 실행이 안될 수 있음

#### All
- __(_Option_)__ [Docker Hub](https://hub.docker.com/)에 가입하고, `docker images`를 실행하여 앞서 만든 이미지들을 push 해야함
- Tagging docker images
    ```bash
    # Anywhere
    (OS)$ docker tag goregtrain:linear <DockerHub ID>/goregtrain:linear
    (OS)$ docker tag goregtrain:multiple <DockerHub ID>/goregtrain:multiple
    (OS)$ docker tag goregpredict:latest <DockerHub ID>/goregpredict:latest

    # Example
    (OS)$ docker tag goregpredict:latest gzupark/goregpredict:latest
    ```
- Pushing docker images
    ```bash
    # Anywhere
    (OS)$ docker push <DockerHub ID>/goregtrain:linear
    (OS)$ docker push <DockerHub ID>/goregtrain:multiple
    (OS)$ docker push <DockerHub ID>/goregpredict:latest

    # Example
    (OS)$ docker push gzupark/goregpredict:latest
    ```
    - Docker Hub에 이미지를 올리기 어렵다면, 아래의 이미지들을 사용
        - [goregtrain](https://hub.docker.com/repository/docker/gzupark/goregtrain)
        - [goregpredict](https://hub.docker.com/repository/docker/gzupark/goregpredict)

#### MacOS
- __(중요)__ 만약 MacOS BigSur로 운영체제를 업데이트했다면, MacOS CTL을 재설치 [issue](https://apple.stackexchange.com/a/406529)
- Minikube [설치](https://minikube.sigs.k8s.io/docs/start/)
    ```bash
    (OS)$ brew install minikube
    ```
- Minikube 실행 및 확인
    ```bash
    (OS)$ minikube start
    # When you finish starting precess
    (OS)$ docker ps  # check running containers
    ```
- kubectl 설치
    ```bash
    (OS)$ minikube kubectl -- get po -A
    ```
- Pachyderm `1.11.x` [설치](https://www.pachyderm.com/getting-started/)
    ```bash
    (OS)$ brew tap pachyderm/tap && brew install pachyderm/tap/pachctl@1.11
    ```
- Pachyderm 확인
    ```bash
    (OS)$ pachctl version --client-only
    # 1.11.7
    ```

#### Windows or Debian-based Linux
WIP

### How to use
[Run with Go files](#run-with-go-files)에서 소개하는 것은 `.go` 파일을 빌드하여 실행하는 것이고, `pachctl`을 활용한 방법은 [pachctl](./pachctl/README.md) 디렉토리를 참고

#### Deploy
- pachyderm 배포
    ```bash
    (OS)$ pachctl deploy local

    # Stay until running status of etcd-xxx & pachd-xxx 
    (OS)$ kubectl get pods

    # NAME                     READY   STATUS              RESTARTS   AGE
    # dash-56fd8c65bf-bdfpr    0/2     ContainerCreating   0          4s
    # etcd-58c9bf64b8-pqqjt    0/1     ContainerCreating   0          5s
    # pachd-78bf7577f6-kwm5r   0/1     ContainerCreating   0          4s

    # NAME                     READY   STATUS              RESTARTS   AGE
    # dash-56fd8c65bf-bdfpr    0/2     ContainerCreating   0          68s
    # etcd-58c9bf64b8-pqqjt    1/1     Running             0          68s
    # pachd-78bf7577f6-kwm5r   1/1     Running             0          68s
    ```
- 포트포워딩
    ```bash
    # Recommend
    (OS)$ pachctl port-forward & pachctl version
    # (option) Connect to directly minikue instance
    (OS)$ pachctl config update context `pachctl config get active-context` \
            --pachd-address=`minikube ip`:30650

    # check configuration
    (OS)$ cat ~/.pachyderm/config.json
    ```

#### Run with Go files
- Create repositories
    ```bash
    # build
    (OS)$ make compile GOOS=darwin GOFILE=01_create_repository.go

    # run
    (OS)$ ./01_create_repository
    ```
- Put files to repositories
    ```bash
    # build
    (OS)$ make compile GOOS=darwin GOFILE=02_put_file.go

    # run
    (OS)$ ./02_put_file
    ```
- Create & update model pipeline
    ```bash
    # build
    (OS)$ make compile GOOS=darwin GOFILE=03_create_model_pipeline.go

    # run
    (OS)$ ./03_create_model_pipeline -tag <linear or multiple> -user <Docker Hub ID>

    # Update
    (OS)$ ./03_create_model_pipeline -tag <linear or multiple> -user <Docker Hub ID>

    # example
    (OS)$ ./03_create_model_pipeline -tag linear -user gzupark
    (OS)$ ./03_create_model_pipeline -tag multiple -user gzupark
    ```
    - Example to check status
        ```bash
        (OS)$ kubectl get pods
        # NAME                      READY   STATUS            RESTARTS   AGE
        # dash-56fd8c65bf-bdfpr     2/2     Running           0          4m
        # etcd-58c9bf64b8-pqqjt     1/1     Running           0          4m1s
        # pachd-78bf7577f6-kwm5r    1/1     Running           0          4m
        # pipeline-model-v1-qdtw5   0/2     PodInitializing   0          21s

        (OS)$ kubectl get pods
        # NAME                      READY   STATUS    RESTARTS   AGE
        # dash-56fd8c65bf-bdfpr     2/2     Running   0          4m19s
        # etcd-58c9bf64b8-pqqjt     1/1     Running   0          4m20s
        # pachd-78bf7577f6-kwm5r    1/1     Running   0          4m19s
        # pipeline-model-v1-qdtw5   2/2     Running   0          40s

        (OS)$ pachctl list job
        # ID    PIPELINE STARTED        DURATION           RESTART PROGRESS  DL       UL   STATE   
        # 8c368 model    34 seconds ago Less than a second 0       1 + 0 / 1 36.72KiB 251B success
        ```
- Create prediction pipeline
    ```bash
    # build
    (OS)$ make compile GOOS=darwin GOFILE=04_create_predict_pipeline.go

    # run
    (OS)$ ./04_create_predict_pipeline -user <Docker Hub ID>

    # example
    (OS)$ ./04_create_predict_pipeline -user gzupark
    ```
    - Example to check status
        ```bash
        (OS)$ kubectl get pods
        # NAME                           READY   STATUS            RESTARTS   AGE
        # dash-56fd8c65bf-bdfpr          2/2     Running           0          6m12s
        # etcd-58c9bf64b8-pqqjt          1/1     Running           0          6m13s
        # pachd-78bf7577f6-kwm5r         1/1     Running           0          6m12s
        # pipeline-model-v1-qdtw5        2/2     Running           0          2m33s
        # pipeline-prediction-v1-9lcsl   0/2     PodInitializing   0          5s

        (OS)$ kubectl get pods
        # NAME                           READY   STATUS    RESTARTS   AGE
        # dash-56fd8c65bf-bdfpr          2/2     Running   0          6m30s
        # etcd-58c9bf64b8-pqqjt          1/1     Running   0          6m31s
        # pachd-78bf7577f6-kwm5r         1/1     Running   0          6m30s
        # pipeline-model-v1-qdtw5        2/2     Running   0          2m51s
        # pipeline-prediction-v1-9lcsl   2/2     Running   0          23s

        (OS)$ pachctl list job
        # ID    PIPELINE   STARTED        DURATION           RESTART PROGRESS  DL       UL   STATE   
        # 353ce prediction 31 seconds ago 1 second           0       1 + 0 / 1 883B     802B success 
        # 8c368 model      2 minutes ago  Less than a second 0       1 + 0 / 1 36.72KiB 251B success 
        ```

#### Useful `pachctl` command
```bash
# Check repositories
(OS)$ pachctl list repo

# Check repository' branch
(OS)$ pachctl inspect repo <repo Name> -v

# Show all commits in the repository
(OS)$ pachctl list commit <repo Name>@<branch>

# Show files commit history in the repository
(OS)$ pachctl list file <repo Name>@<branch> --history all

# Print out the file
(OS)$ pachctl get file <repo Name>@<branch>:<file Name>

# Check finished job
(OS)$ pachctl list job
# Show logs of job via ID
(OS)$ pachctl logs --job=<ID>
```

#### Undeploy
```bash
(OS)$ pachctl undeploy
# Very important to disconnect port forwarding
(OS)$ kill -9 $(ps u | grep 'pachctl port-forward' | grep -v color | awk '{print $2}')
```

#### Clean up
```bash
(OS)$ pachctl delete all

# Run to undeploy process
(OS)$ pachctl undeploy
(OS)$ kill -9 $(ps u | grep 'pachctl port-forward' | grep -v color | awk '{print $2}')

(OS)$ rm ~/.pachyderm/config.json
(OS)$ minikube delete
```
