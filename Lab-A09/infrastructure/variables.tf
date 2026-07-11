variable "resource_group_name" {
  description = "Name of the Azure resource group."
  type        = string
  default     = "rg-cst8918-hybrid-a09"
}

variable "location" {
  description = "Azure region where the resource group will be created."
  type        = string
  default     = "Canada Central"
}