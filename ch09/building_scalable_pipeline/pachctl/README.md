## Run with `pachctl`
Pachyderm이 설치된 local OS에서 command line tool을 활용한 방법

#### How to use

- Create repositories
    ```bash
    # Anywhere
    (os)$ pachctl create repo training --description "Diabetes training data (.csv)"
    (os)$ pachctl create repo attributes --description "Attributes data (.json)"
    ```
- Put files to repositories
    ```bash
    # Move to file located
    (os)$ pachctl put file training@master --overwrite --file diabetes_training.csv

    # Move to project directory (ex: ${HOME}/ml_with_go)
    (os)$ pachctl put file attributes@master:/ --recursive --overwrite --file storage/attributes
    ```
- Create & update model pipeline
    - Create
        ```bash
        # Move to this file located in
        (os)$ pachctl create pipeline --file model_linear.json 
        ```
    - Update
        ```bash
        (OS)$ pachctl update pipeline --reprocess --file model_multiple.json 
        ```
- Create prediction pipeline
    ```bash
    # Move to this file located in
    (os)$ pachctl create pipeline --file prediction.json 
    ```