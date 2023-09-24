#!/bin/bash

set -euo pipefail

# Create Public S3 Bucket
awslocal s3api create-bucket \
  --bucket lgtmgen-local-images \
  --acl public-read
