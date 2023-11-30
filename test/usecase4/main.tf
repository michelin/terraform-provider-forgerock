terraform {
  required_providers {
    forgerock = {
      source = "terraform.local/local/forgerock"
    }
  }
  backend "local" {
    path = "../terraform-loc.tfstate"
  }
}

provider "forgerock" {
  username      = var.username
  password      = var.password
  forgerock_api = var.forgerock_api
  realm_path    = var.realm_path
  mail_sender = {
    send_client_secret_mail = true
    smtp_server             = "smtp.example.com"
    smtp_port               = 587
    sender_email            = "sender@example.com"
    sender_username         = "username"
    sender_password         = "password"
  }
}