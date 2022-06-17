data "yandex_container_registry" "main" {
  name      = "main"
  folder_id = data.yandex_resourcemanager_folder.mai-backend.id
}

resource "yandex_iam_service_account" "puller-sa" {
  folder_id = data.yandex_resourcemanager_folder.mai-backend.id
  name      = "puller-sa"
}

resource "yandex_resourcemanager_folder_iam_member" "puller-sa-puller" {
  folder_id = data.yandex_resourcemanager_folder.mai-backend.id
  role      = "container-registry.images.puller"
  member    = "serviceAccount:${yandex_iam_service_account.puller-sa.id}"
}
