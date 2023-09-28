output "images_url" {
  value = "https://${module.cloudfront_images.this.domain_name}"
}
