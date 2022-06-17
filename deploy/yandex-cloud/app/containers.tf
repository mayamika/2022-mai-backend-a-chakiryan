locals {
  api_server_image = "api-server:0.1.0"
  ui_server_image  = "ui-server:0.1.0"
}

resource "yandex_serverless_container" "api-server" {
  name               = "api-server"
  memory             = 256
  execution_timeout  = "15s"
  cores              = 1
  core_fraction      = 5
  service_account_id = yandex_iam_service_account.puller-sa.id

  image {
    url = "cr.yandex/${data.yandex_container_registry.main.id}/${local.api_server_image}"
    environment = {
      POSTGRES             = "postgres://${yandex_mdb_postgresql_user.api-server.name}:${yandex_mdb_postgresql_user.api-server.password}@c-${data.yandex_mdb_postgresql_cluster.main.id}.rw.mdb.yandexcloud.net:6432/${yandex_mdb_postgresql_database.api-server.name}?sslmode=verify-full&sslrootcert=/certs/root.crt"
      OPENSEARCH_ADDRESSES = "https://c-${data.yandex_mdb_elasticsearch_cluster.main.id}.rw.mdb.yandexcloud.net:9200/"
      OPENSEARCH_USERNAME  = "admin"
      OPENSEARCH_PASSWORD  = "my-cool-pass"
      OPENSEARCH_CA_CERT   = "/certs/root.crt"
      S3_ADDRESS           = "storage.yandexcloud.net"
      S3_ACCESS_KEY_ID     = yandex_iam_service_account_static_access_key.storage-sa-static-key.access_key
      S3_ACCESS_KEY_SECRET = yandex_iam_service_account_static_access_key.storage-sa-static-key.secret_key
      IMAGES_BUCKET        = yandex_storage_bucket.images.bucket
    }
  }
}

resource "yandex_serverless_container" "ui-server" {
  name               = "ui-server"
  memory             = 256
  execution_timeout  = "15s"
  cores              = 1
  core_fraction      = 5
  service_account_id = yandex_iam_service_account.puller-sa.id

  image {
    url = "cr.yandex/${data.yandex_container_registry.main.id}/${local.ui_server_image}"
    environment = {
      API_SERVER_ADDR = trimsuffix(yandex_serverless_container.api-server.url, "/")
      IMAGES_ADDR     = "https://${yandex_storage_bucket.images.bucket}.storage.yandexcloud.net/"
    }
  }
}
