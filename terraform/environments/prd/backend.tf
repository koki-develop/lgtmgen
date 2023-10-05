terraform {
  backend "s3" {
    region  = "us-east-1"
    bucket  = "lgtmgen-tfstates"
    encrypt = true
    key     = "prd/terraform.tfstate"
  }
}
