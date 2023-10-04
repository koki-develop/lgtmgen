#!/bin/bash

set -euo pipefail

readonly RESOURCE_PREFIX="lgtmgen-local"

# Create S3 Buckets
awslocal s3api create-bucket \
  --bucket ${RESOURCE_PREFIX}-images \
  --acl public-read

awslocal s3api create-bucket \
  --bucket ${RESOURCE_PREFIX}-news \
  --acl private

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

awslocal dynamodb create-table \
  --table-name ${RESOURCE_PREFIX}-reports \
  --key-schema \
    AttributeName=id,KeyType=HASH \
  --attribute-definitions \
    AttributeName=id,AttributeType=S \
  --billing-mode PAY_PER_REQUEST

awslocal dynamodb create-table \
  --table-name ${RESOURCE_PREFIX}-rates \
  --key-schema \
    AttributeName=ip,KeyType=HASH \
    AttributeName=tier,KeyType=RANGE \
  --attribute-definitions \
    AttributeName=ip,AttributeType=S \
    AttributeName=tier,AttributeType=S \
  --billing-mode PAY_PER_REQUEST

# Create SQS Queues
awslocal sqs create-queue \
  --queue-name ${RESOURCE_PREFIX}-notifications
