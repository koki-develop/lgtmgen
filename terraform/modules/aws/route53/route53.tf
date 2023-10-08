data "aws_route53_zone" "main" {
  name = var.domains.apex
}

resource "aws_route53_record" "ui_apex" {
  count = var.domains.ui != null ? 1 : 0

  zone_id = data.aws_route53_zone.main.zone_id
  name    = var.domains.apex
  type    = "A"
  records = ["76.76.21.21"] # Vercel
  ttl     = 60
}

resource "aws_route53_record" "ui_www" {
  count = var.domains.ui != null ? 1 : 0

  zone_id = data.aws_route53_zone.main.zone_id
  name    = var.domains.ui
  type    = "CNAME"
  records = ["cname.vercel-dns.com"]
  ttl     = 60
}

resource "aws_route53_record" "api" {
  zone_id = data.aws_route53_zone.main.id
  name    = var.domains.api
  type    = "A"

  alias {
    name                   = var.routings.api.domain_name
    zone_id                = var.routings.api.zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "images" {
  zone_id = data.aws_route53_zone.main.id
  name    = var.domains.images
  type    = "A"

  alias {
    name                   = var.routings.images.domain_name
    zone_id                = var.routings.images.zone_id
    evaluate_target_health = false
  }
}
