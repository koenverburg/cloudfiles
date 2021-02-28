provider "azurerm" {
  version = "~>2.0"
  features {}
}

resource "azurerm_resource_group" "rg" {
  name     = "tf-rg-linux-vm"
  location = "westeurope"

  tags = {
    environment = "Terraform Demo"
  }
}

resource "azurerm_virtual_network" "vnet" {
  name                = "vnet-linux-vm"
  address_space       = ["10.0.0.0/16"]
  location            = "westeurope"
  resource_group_name = azurerm_resource_group.rg.name

  tags = {
    environment = "Terraform Demo"
  }
}

resource "azurerm_subnet" "subnet" {
  name                 = "vnet-linux-vm"
  resource_group_name  = azurerm_resource_group.rg.name
  virtual_network_name = azurerm_virtual_network.vnet.name
  address_prefixes     = ["10.0.2.0/24"]
}
