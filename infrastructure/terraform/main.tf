provider "aws" {
  region = "us-west-2"
}

resource "aws_s3_bucket" "fraud_detection_data" {
  bucket = "fraud-detection-data"
  acl    = "private"
}

resource "aws_eks_cluster" "fraud_detection_cluster" {
  name     = "fraud-detection"
  role_arn = "arn:aws:iam::123456789012:role/EKSRole"

  vpc_config {
    subnet_ids = ["subnet-abc123", "subnet-def456"]
  }
}