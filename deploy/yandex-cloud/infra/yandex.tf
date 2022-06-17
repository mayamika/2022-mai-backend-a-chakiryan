terraform {
  required_providers {
    yandex = {
      source = "yandex-cloud/yandex"
    }
  }
  required_version = ">= 0.13"

  backend "s3" {
    endpoint = "storage.yandexcloud.net"
    bucket   = "backend-tfstate"
    region   = "ru-central1"
    key      = "infra.tfstate"

    skip_region_validation      = true
    skip_credentials_validation = true
  }
}

provider "yandex" {
  zone = "ru-central1-a"
}

data "yandex_resourcemanager_folder" "mai-backend" {
  name = "mai-backend"
}
