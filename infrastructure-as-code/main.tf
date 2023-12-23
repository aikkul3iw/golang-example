provider "aws" {
  region = "ap-southeast-1"
}

module "eks" {
  source = "terraform-aws-modules/eks/aws"

  cluster_name = "my-eks-cluster"
  subnets      = ["subnet-xxxxxxxxxxxxxxxxx", "subnet-yyyyyyyyyyyyyyyyy"]

  cluster_create_timeout = "30m"
}

resource "aws_iam_role" "eks_worker_nodes" {
  name = "eks-worker-nodes-role"

  assume_role_policy = <<EOF
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "ec2.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }
  EOF
}

resource "aws_iam_role_policy_attachment" "eks_s3_policy_attachment" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
  role       = aws_iam_role.eks_worker_nodes.name
}

resource "aws_iam_role_policy_attachment" "eks_sqs_policy_attachment" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonSQSFullAccess"
  role       = aws_iam_role.eks_worker_nodes.name
}