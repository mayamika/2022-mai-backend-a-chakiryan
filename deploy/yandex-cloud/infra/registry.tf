resource "yandex_container_registry" "main" {
  name      = "main"
  folder_id = data.yandex_resourcemanager_folder.mai-backend.id
}

output "registry_id" {
  value = yandex_container_registry.main.id
}
