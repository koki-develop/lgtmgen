output "certificate_api_arn" {
  value = aws_acm_certificate.api.arn
}

output "certificate_images_arn" {
  value = aws_acm_certificate.images.arn
}
