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

module "s3_original_images" {
  source = "../../modules/aws/s3"
  name   = local.name
  tier   = "original-images"
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
  certificate_arn           = module.route53.certificate_images_arn
}

module "ecr" {
  source = "../../modules/aws/ecr"
  name   = local.name
}

module "route53" {
  source = "../../modules/aws/route53"

  domains = {
    apex   = local.domain
    api    = local.domain_api
    images = local.domain_images
  }

  routings = {
    api = {
      domain_name = module.api_gateway.domain_name.cloudfront_domain_name
      zone_id     = module.api_gateway.domain_name.cloudfront_zone_id
    }
    images = {
      domain_name = module.cloudfront_images.this.domain_name
      zone_id     = module.cloudfront_images.this.hosted_zone_id
    }
  }
}

module "api_gateway" {
  source          = "../../modules/aws/api_gateway"
  stage           = local.stage
  api_name        = local.name
  domain          = local.domain_api
  certificate_arn = module.route53.certificate_api_arn
}
