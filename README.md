# Terraform-provider-forgerock

[![GitHub Build](https://img.shields.io/github/actions/workflow/status/michelin/terraform-provider-forgerock/on_push_master.yml?branch=master&logo=github&style=for-the-badge)](https://img.shields.io/github/actions/workflow/status/michelin/terraform-provider-forgerock/on_push_master.yml)
[![GitHub release](https://img.shields.io/github/v/release/michelin/terraform-provider-forgerock?logo=github&style=for-the-badge)](https://github.com/michelin/terraform-provider-forgerock/releases)
[![GitHub commits since latest release (by SemVer)](https://img.shields.io/github/commits-since/michelin/terraform-provider-forgerock/latest?logo=github&style=for-the-badge)](https://github.com/michelin/terraform-provider-forgerock/commits/main)
[![GitHub Stars](https://img.shields.io/github/stars/michelin/terraform-provider-forgerock?logo=github&style=for-the-badge)](https://github.com/michelin/terraform-provider-forgerock)
[![GitHub Watch](https://img.shields.io/github/watchers/michelin/terraform-provider-forgerock?logo=github&style=for-the-badge)](https://github.com/michelin/terraform-provider-forgerock)
[![SonarCloud Coverage](https://img.shields.io/sonar/coverage/michelin_terraform-provider-forgerock?logo=sonarcloud&server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge)](https://sonarcloud.io/component_measures?id=michelin_terraform-provider-forgerock&metric=coverage&view=list)
[![SonarCloud Tests](https://img.shields.io/sonar/tests/michelin_terraform-provider-forgerock/master?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonarcloud)](https://sonarcloud.io/component_measures?metric=tests&view=list&id=michelin_kstreamplify)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?logo=apache&style=for-the-badge)](https://opensource.org/licenses/Apache-2.0)

terraform-provider-forgerock introduces Forgerock OAuth2 client creation functionality to terraform.

## Table of Contents

* [Principles](#principles)
* [Local run](#local-run)
* [Build project](#build-project)
* [Provider configuration](#provider-configuration)
* [Resource configuration](#resource-configuration)
  * [Public client (Front to Back)](#public-client-front-to-back)
  * [Private client (Back to Back)](#private-client-back-to-back)
  * [Resource complet field list](#resource-complet-field-list)
* [Example](#example)
* [Contribution](#contribution)

## Principles

Terraform-provider-forgerock is a terraform provider that allows you to create Forgerock OAuth2 clients through ForgeRock APIs.

## Local run

To start the provider in debug mode, you can use Visual Studio Code:

* Navigate to the `/example` directory.
* Create a `terraform.tfvars` file (do not track in Git) and fill it out.
* Press F5.
* Execute the command provided in the console after the provider starts.

## Build project

To build your project run these commands:

```bash
go mod tidy
go build
```

## Provider configuration

To configure the provider you need to add the following code to your terraform file:

```hcl
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
```

## Resource configuration

We provide a set of default configurations for several types of clients:

* Public client (authentication code flow)
* Private client (client secret)

### Public client (code flow)

```hcl
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
```

### Private client (client secret)

```hcl
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
```

### Example

You can find a complete example [here](./example/main.tf)

### Resource complete field list

If you want to customize the default configuration given above you can refer to the following [documentation](./docs/fields_list.md)
