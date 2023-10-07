resource "aws_ecr_repository" "base" {
  name                 = "${var.name}-api-base"
  image_tag_mutability = "MUTABLE"
}
