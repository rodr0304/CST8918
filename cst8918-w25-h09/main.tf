resource "azurerm_resource_group" "rg" {
  name     = "rg-cst8918-h09"
  location = "Canada Central"
}

resource "azurerm_kubernetes_cluster" "app" {

  name                = "aks-cst8918"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name

  dns_prefix = "cst8918"

  default_node_pool {

    name = "default"

    vm_size = "Standard_B2s"

    auto_scaling_enabled = true

    min_count = 1
    max_count = 3
  }

  identity {
    type = "SystemAssigned"
  }
}