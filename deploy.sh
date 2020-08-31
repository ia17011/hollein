#!/bin/bash

sam build 

sam package \
    --output-template-file template.yaml \
    --s3_bucket $S3_BUCKET

sam deploy --guided \
    --region ap-northeast-1 \
    --template-file template.yaml \
    --stack-name app \
    --s3_bucket $S3_BUCKET \
    --capabilities CAPABILITY_IAM