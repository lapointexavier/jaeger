resource "aws_dynamodb_table" "spanstore_services_operations" {
  name           = "spanstore_services_operations"
  read_capacity  = "10"
  write_capacity = "10"
  hash_key       = "service"
  range_key      = "operation"

  attribute {
    name = "service"
    type = "S"
  }

  attribute {
    name = "operation"
    type = "S"
  }
}
