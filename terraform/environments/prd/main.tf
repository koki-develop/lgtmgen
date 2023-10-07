data "aws_route53_zone" "main" {
  name = local.domain
}

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

module "s3_news" {
  source = "../../modules/aws/s3"
  name   = local.name
  tier   = "news"
}

module "cloudfront_images" {
  source                    = "../../modules/aws/cloudfront"
  name                      = local.name
  tier                      = "images"
  origin_bucket_domain_name = module.s3_images.this.bucket_regional_domain_name
  domain                    = local.domain_images
  certificate_arn           = aws_acm_certificate.images.arn
}

module "ecr" {
  source = "../../modules/aws/ecr"
  name   = local.name
}
