locals {
  stage = "prd"
  name  = "lgtmgen-${local.stage}"

  domain        = "lgtmgen.org"
  domain_api    = "api.${local.domain}"
  domain_images = "images.${local.domain}"
}
