data "aws_api_gateway_rest_api" "main" {
  name = var.api_name
}

resource "aws_api_gateway_domain_name" "main" {
  certificate_arn = var.certificate_arn
  domain_name     = var.domain
  security_policy = "TLS_1_2"
}

resource "aws_api_gateway_base_path_mapping" "main" {
  api_id      = data.aws_api_gateway_rest_api.main.id
  stage_name  = var.stage
  domain_name = aws_api_gateway_domain_name.main.domain_name
}
