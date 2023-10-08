resource "aws_acm_certificate" "api" {
  domain_name       = var.domains.api
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate" "images" {
  domain_name       = var.domains.images
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}
