#!/bin/bash

URL="https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.3.0.tar.gz"

curl -L $URL | tar -C "/usr/local" -xz
ldconfig

gcc hello_tf.c -ltensorflow -o hello_tf && \
    ./hello_tf
