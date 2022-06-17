resource "yandex_mdb_elasticsearch_cluster" "main" {
  name        = "main"
  environment = "PRESTABLE"
  network_id  = yandex_vpc_network.main.id

  config {
    admin_password = "my-cool-pass"

    data_node {
      resources {
        resource_preset_id = "s2.micro"
        disk_type_id       = "network-ssd"
        disk_size          = 10
      }
    }
  }

  host {
    name             = "node"
    zone             = "ru-central1-a"
    type             = "DATA_NODE"
    assign_public_ip = true
    subnet_id        = yandex_vpc_subnet.main.id
  }

  maintenance_window {
    type = "ANYTIME"
  }
}
