## Image Classification with TensorFlow Go
[01_tf_image_classification.go](./01_tf_image_classification.go) 파일은 [TF Go example](https://github.com/tensorflow/tensorflow/blob/0d1e4cf5b7dc60b1f0a45eb06a120df058ff4077/tensorflow/go/example_inception_inference_test.go)과 동일

#### How to use
```bash
# build
(docker)$ go build 01_tf_image_classification.go
# run
(docker)$ ./01_tf_image_classification -image=$MLGO/storage/data/airplane.jpg
(docker)$ ./01_tf_image_classification -image=$MLGO/storage/data/gopher.jpg
(docker)$ ./01_tf_image_classification -image=$MLGO/storage/data/pug.jpg
```
