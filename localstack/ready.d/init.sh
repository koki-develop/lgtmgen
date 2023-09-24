#!/bin/bash

set -euo pipefail

readonly RESOURCE_PREFIX="lgtmgen-local"

# Create Public S3 Bucket
awslocal s3api create-bucket \
  --bucket ${RESOURCE_PREFIX}-images \
  --acl public-read

# Create DynamoDB Tables
awslocal dynamodb create-table \
  --table-name ${RESOURCE_PREFIX}-lgtms \
  --key-schema \
    AttributeName=id,KeyType=HASH \
    AttributeName=created_at,KeyType=RANGE \
  --attribute-definitions \
    AttributeName=id,AttributeType=S \
    AttributeName=created_at,AttributeType=S \
    AttributeName=status,AttributeType=S \
  --global-secondary-indexes \
    IndexName=index_by_status,KeySchema=["{AttributeName=status,KeyType=HASH}","{AttributeName=created_at,KeyType=RANGE}"],Projection="{ProjectionType=ALL}" \
  --billing-mode PAY_PER_REQUEST
