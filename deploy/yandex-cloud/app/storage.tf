resource "yandex_iam_service_account" "storage-sa" {
  folder_id = data.yandex_resourcemanager_folder.mai-backend.id
  name      = "storage-sa"
}

resource "yandex_resourcemanager_folder_iam_member" "storage-sa-editor" {
  folder_id = data.yandex_resourcemanager_folder.mai-backend.id
  role      = "storage.editor"
  member    = "serviceAccount:${yandex_iam_service_account.storage-sa.id}"
}

resource "yandex_resourcemanager_folder_iam_member" "storage-sa-admin" {
  folder_id = data.yandex_resourcemanager_folder.mai-backend.id
  role      = "storage.admin"
  member    = "serviceAccount:${yandex_iam_service_account.storage-sa.id}"
}

resource "yandex_iam_service_account_static_access_key" "storage-sa-static-key" {
  service_account_id = yandex_iam_service_account.storage-sa.id
  description        = "static access key for object storage"
}

resource "yandex_storage_bucket" "images" {
  bucket                = "mai-backend-images"
  default_storage_class = "COLD"
  acl                   = "public-read"
  access_key            = yandex_iam_service_account_static_access_key.storage-sa-static-key.access_key
  secret_key            = yandex_iam_service_account_static_access_key.storage-sa-static-key.secret_key
  force_destroy         = true
}
