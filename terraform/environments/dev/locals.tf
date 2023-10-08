locals {
  stage = "dev"
  name  = "lgtmgen-${local.stage}"

  domain        = "lgtmgen.org"
  domain_api    = "dev.api.${local.domain}"
  domain_images = "dev.images.${local.domain}"
}
