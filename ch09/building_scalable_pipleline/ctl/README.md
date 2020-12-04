## Pachyderm CTL
Pachyderm이 설치된 local OS에서 command line tool을 활용한 방법

#### Steps

1. Repository 생성
    ```bash
    # Anywhere
    (os)$ pachctl create repo training --description "Diabetes training data (.csv)"
    (os)$ pachctl create repo attributes --description "Attributes data (.json)"
    ```
2. 해당 repository에 file 올리기
    ```bash
    # Move to file located
    (os)$ pachctl put file training@master --overwrite --file diabetes_training.csv
    # Move to project directory (ex: ${HOME}/ml_with_go)
    (os)$ pachctl put file attributes@master:/ --recursive --overwrite --file storage/attributes
    ```
3. Pipeline 생성
    - Create
    ```bash
    # Move to this file located in
    (os)$ pachctl create pipeline --file model_linear.json 
    ```
    - Update
    ```bash
    (OS)$ pachctl update pipeline --reprocess --file model_multiple.json 
    ```
