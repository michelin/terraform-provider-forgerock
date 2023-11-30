terraform {
  required_providers {
    forgerock = {
      source = "michelin/forgerock"
    }
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

resource "forgerock_oauth2Client" "myPublicClient" {

  name = "my_public_client"
  admin_mail = "yourmail@mail.com"

  advanced_oauth2_client_config = {
    token_endpoint_auth_method = "none"
    grant_types                = ["authorization_code", "refresh_token"]
    is_consent_implied = true
  }

  core_open_id_client_config = {
    post_logout_redirect_uri = ["http://localhost:4200"]
  }

  core_oauth2_client_config = {
    status = "Active"
    scopes = ["profile", "email", "openid"]
    redirection_uris = ["http://localhost:4200", "https://anotherurl.com"]
    client_type = "Public"
  }
}


resource "forgerock_oauth2Client" "myPrivateClient" {

  name = "my_private_client"
  admin_mail = "yourmail@mail.com"
  user_password_version = 0

  advanced_oauth2_client_config = {
    token_endpoint_auth_method = "none"
    grant_types = ["client_credentials"]
    is_consent_implied = true
  }

  core_open_id_client_config = {
    post_logout_redirect_uri = [""]
  }

  core_oauth2_client_config = {
    status = "Active"
    scopes = ["profile", "email", "openid"]
    redirection_uris = [""]
    client_type = "Confidential"
  }
}