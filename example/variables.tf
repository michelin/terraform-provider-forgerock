variable "username" {
  type = string
}

variable "password" {
  type = string
}

variable "forgerock_api" {
  type        = string
  description = "The URL of the ForgeRock API"
}

variable "realm_path" {
  type        = string
  description = "The path to the realm"
}

variable "sas_token" {
  type        = string
  description = "The SAS token to access the storage account"
}
