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

resource "forgerock_oauth2Client" "INTEGRATION_TEST_AUTO_CLIENT_1" {
  name = "INTEGRATION_TEST_AUTO_CLIENT_1"
  admin_mail = "INTEGRATION_TEST_AUTO_CLIENT_1@michelin.com"
  user_password_version = 1
}

resource "forgerock_oauth2Client" "INTEGRATION_TEST_AUTO_CLIENT_2" {
  name = "INTEGRATION_TEST_AUTO_CLIENT_2"
  admin_mail = "INTEGRATION_TEST_AUTO_CLIENT_2@michelin.com"
  user_password_version = 1

  core_oauth2_client_config = {
    status      = "Active"
    scopes      = ["profile","openid","test2"]
    client_type = "Confidential"
  }

  advanced_oauth2_client_config = {
    token_endpoint_auth_method = "private_key_jwt"
    grant_types                = ["client_credentials"]
  }

  sign_enc_oauth2_client_config = {
    jwk_set             = "{   \"keys\": [       {           \"kty\": \"RSA\",           \"e\": \"AQAB\",            \"use\": \"sig\",           \"kid\": \"23c50663-6902-4cdb-9ef9-a15f9bb10eeb\",          \"n\": \"xsczx0C6tMxxZquC-bsAfCd7HGx41c8aje5pn5Oys6lSUKEJ7mqZFWQi86IkOWwjoorSpczP1xOKhwN5_80_Yi_zAYs7iDeENDXt-O5bjNdagC3nxgGYoefSaJgmKmK3Da20b_YcIWGEddS_IK4QRtgLEcY3wh6-9fUvHsbCSarPGdm34E4F1jAaiuC1dTyT5qUiDroiK8qig27iiIOHXGUz2TpSrpHB5bWvTP6nELLN2m05dG5gF0EA8H3WjCfMrVPM11avgLt5TOOKNR8u5lQZvNVoUY_f1X_cUfhRyNuTpnJY6WOVQy-lbG0XQp4Wbuske3-6hlAW_JSFIsySWw\"       }   ] }"
    public_key_location = "jwks"
  }

}

resource "forgerock_oauth2Client" "INTEGRATION_TEST_AUTO_CLIENT_FAIL_3_1" {
  name = "INTEGRATION_TEST_AUTO_CLIENT_3"
  admin_mail = "INTEGRATION_TEST_AUTO_CLIENT_2@michelin.com"
}
