resource "azurerm_resource_group" "example" {
  name=var.resource_group_name
  location = var.location

  tags = {
    Environment = "Development"
    Assignment  = "Hybrid-A09"
    ManagedBy   = "Terraform"
  }
}
