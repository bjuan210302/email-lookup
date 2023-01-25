terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_default_vpc" "default" {}

data "http" "my_ip" {
  url = "http://ipv4.icanhazip.com"
}

resource "aws_security_group" "ec2_sg" {
  name        = "elookup-security-group"
  description = "Allow hhtp access on port 6002 for backend"
  vpc_id      = aws_default_vpc.default.id

  ingress {
    description = "http access"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "ssh access"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["${chomp(data.http.my_ip.body)}/32"]
  }

  ingress {
    description = "zinc-server access"
    from_port   = 6001
    to_port     = 6001
    protocol    = "tcp"
    cidr_blocks = ["${chomp(data.http.my_ip.body)}/32"]
  }

  ingress {
    description = "e-lookup-be access"
    from_port   = 6002
    to_port     = 6002
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "e-lookup-fe access"
    from_port   = 6003
    to_port     = 6003
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    project = "email-lookup"
  }
}

resource "aws_iam_instance_profile" "ec2_deployer_user" {
  name = "terraform_user"
}

data "aws_ami" "amazon_linux_2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-ebs"]
  }
}

resource "aws_instance" "elookup_ec2" {
  ami                    = data.aws_ami.amazon_linux_2.id
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.ec2_sg.id]
  iam_instance_profile   = aws_iam_instance_profile.ec2_deployer_user.name
  key_name               = "terraform-key-pair"

  user_data = <<-EOF
    #!/bin/bash
    set -ex
    sudo yum update -y
    sudo amazon-linux-extras install docker -y
    sudo usermod -a -G docker ec2-user
    sudo systemctl start docker
    sudo yum install git -y
    sudo curl -L https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose

    # clone repo and install ZincSearch
    git clone https://github.com/bjuan210302/email-lookup.git /home/ec2-user/email-lookup
    sudo curl -L https://github.com/zinclabs/zinc/releases/download/v0.3.6/zinc_0.3.6_Linux_x86_64.tar.gz -o /home/ec2-user/email-lookup/zinc-server/zinc.tar.gz
    sudo tar -xf /home/ec2-user/email-lookup/zinc-server/zinc.tar.gz -C /home/ec2-user/email-lookup/zinc-server/

    public_ip=$(curl http://169.254.169.254/latest/meta-data/public-hostname)
    export ELOOKUPBE_REMOTE_HOST=http://$public_ip:6002/

    cd /home/ec2-user/email-lookup
    docker-compose up
  EOF

  tags = {
    project = "email-lookup"
  }
}
