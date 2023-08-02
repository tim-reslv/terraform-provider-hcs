terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/HuaweiCloudStack"
#      version = "~> 1.26.0"
      version = ">= 1.0.0"
    }
  }
}

# Configure the HuaweiCloud Provider
provider "huaweicloud" {
  region     = "cn-north-4"
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}

# Create a VPC
resource "huaweicloud_vpc" "example" {
  name = "my_vpc"
  cidr = "192.168.0.0/16"
}
