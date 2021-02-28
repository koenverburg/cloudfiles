source "azure-arm" "basic-example" {
  client_id = ""
  client_secret = ""
  subscription_id = ""

  # storage_account = "virtualmachines"
  # resource_group_name = "packertest"
  # capture_container_name = "images"
  # capture_name_prefix = "packer"

  managed_image_name = "ubuntu-18"
  managed_image_resource_group_name = "packertest"

  os_type = "Linux"
  image_publisher = "Canonical"
  image_offer = "UbuntuServer"
  image_sku = "18.04-LTS"

  azure_tags = {
    dept = "engineering"
    task = "Image Development"
  }

  location = "West Europe"
  vm_size = "Standard_DS2_v2"
}

build {
  sources = ["sources.azure-arm.basic-example"]
}