#!/bin/bash

sam build --use-container

sam package \
    --output-template-file template.yaml \
    --s3-bucket $S3_BUCKET

sam deploy --template-file template.yaml \
    --stack-name app \
    --capabilities CAPABILITY_NAMED_IAM CAPABILITY_IAM