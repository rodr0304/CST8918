output "resource_group_name" {
  description = "Name of the created resource group."
  value       = azurerm_resource_group.example.name
}

output "resource_group_location" {
  description = "Location of the created resource group."
  value       = azurerm_resource_group.example.location
}