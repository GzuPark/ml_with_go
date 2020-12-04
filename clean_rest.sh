#!/bin/bash

rm ${PWD}/storage/data/advertising_* 2>/dev/null
rm ${PWD}/storage/data/clean_* 2>/dev/null
rm ${PWD}/storage/data/diff_* 2>/dev/null
rm ${PWD}/storage/data/*diff* 2>/dev/null
rm ${PWD}/storage/data/*db 2>/dev/null
rm ${PWD}/storage/data/*.json 2>/dev/null
rm ${PWD}/storage/model/*.json 2>/dev/null
rm -rf ${PWD}/ch08/deep_learning/model 2>/dev/null
ls ${PWD}/ch08/deep_learning/[0-9][0-9]_*[^go] 2>/dev/null | xargs -d"\n" rm 2>/dev/null
ls ${PWD}/ch09/building_scalable_pipeline/[0-9][0-9]_*[^go] 2>/dev/null | xargs -d"\n" rm 2>/dev/null
ls ${PWD}/ch09/linear_regression/[0-9][0-9]_*[^go] 2>/dev/null | xargs -d"\n" rm 2>/dev/null
ls ${PWD}/ch09/multiple_regression/[0-9][0-9]_*[^go] 2>/dev/null | xargs -d"\n" rm 2>/dev/null
ls ${PWD}/ch09/predict_regression/[0-9][0-9]_*[^go] 2>/dev/null | xargs -d"\n" rm 2>/dev/null
ls ${PWD}/ch09/predict_regression/result/*.json 2>/dev/null | xargs -d"\n" rm 2>/dev/null
