provider "aws" {
  shared_credentials_file = "/Users/xavierlapointe/.aws/credentials"
  profile                 = "dev"
  insecure                = true
  region                  = "us-east-1"

  endpoints {
    dynamodb = "http://localhost:8000"
  }
}
