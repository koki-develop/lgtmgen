module "github_actions" {
  source = "../../modules/aws/github_actions"
  name   = local.name
}

module "s3_images" {
  source = "../../modules/aws/s3"
  name   = local.name
  bucket = "images"
}
