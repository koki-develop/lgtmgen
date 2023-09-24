data "aws_iam_openid_connect_provider" "main" {
  url = "https://token.actions.githubusercontent.com"
}

data "aws_iam_policy_document" "main_assume_role_policy" {
  statement {
    effect  = "Allow"
    actions = ["sts:AssumeRoleWithWebIdentity"]

    principals {
      type        = "Federated"
      identifiers = [data.aws_iam_openid_connect_provider.main.arn]
    }

    condition {
      test     = "StringEquals"
      variable = "token.actions.githubusercontent.com:aud"
      values   = ["sts.amazonaws.com"]
    }

    condition {
      test     = "StringLike"
      variable = "token.actions.githubusercontent.com:sub"
      values   = ["repo:${local.owner}/${local.repo}:*"]
    }
  }
}

resource "aws_iam_role" "main" {
  name               = "${var.name}-github-actions-role"
  assume_role_policy = data.aws_iam_policy_document.main_assume_role_policy.json
}

resource "aws_iam_role_policy_attachment" "main_administrator_access" {
  role       = aws_iam_role.main.name
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}
