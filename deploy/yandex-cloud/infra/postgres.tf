resource "yandex_mdb_postgresql_cluster" "main" {
  name        = "postgres"
  environment = "PRESTABLE"
  network_id  = yandex_vpc_network.main.id

  config {
    version = 14
    resources {
      resource_preset_id = "s2.micro"
      disk_type_id       = "network-ssd"
      disk_size          = 16
    }
  }

  maintenance_window {
    type = "WEEKLY"
    day  = "SAT"
    hour = 12
  }

  host {
    zone             = "ru-central1-a"
    subnet_id        = yandex_vpc_subnet.main.id
    assign_public_ip = true
  }
}
