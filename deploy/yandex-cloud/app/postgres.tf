data "yandex_mdb_postgresql_cluster" "main" {
  name = "postgres"
}

resource "random_password" "api-server-postgres" {
  length  = 16
  special = false
}

resource "yandex_mdb_postgresql_user" "api-server" {
  cluster_id = data.yandex_mdb_postgresql_cluster.main.id
  name       = "api-server"
  password   = random_password.api-server-postgres.result
}

resource "yandex_mdb_postgresql_database" "api-server" {
  cluster_id = data.yandex_mdb_postgresql_cluster.main.id
  name       = "api-server"
  owner      = yandex_mdb_postgresql_user.api-server.name
  lc_collate = "en_US.UTF-8"
  lc_type    = "en_US.UTF-8"

  extension {
    name = "uuid-ossp"
  }
  extension {
    name = "xml2"
  }
}
