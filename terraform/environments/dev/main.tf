module "github_actions" {
  source = "../../modules/aws/github_actions"
  name   = local.name
}

module "s3_images" {
  source           = "../../modules/aws/s3"
  name             = local.name
  tier             = "images"
  distribution_arn = module.cloudfront_images.this.arn
}

module "cloudfront_images" {
  source                    = "../../modules/aws/cloudfront"
  name                      = local.name
  tier                      = "images"
  origin_bucket_domain_name = module.s3_images.this.bucket_regional_domain_name
}
