terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = "us-east1"
}

resource "google_storage_bucket" "static" {
  name                        = "shane-test-bucket"
  location                    = "US"
  storage_class               = "STANDARD"
  force_destroy               = true
  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "default" {
  name         = "example.txt"
  source       = "../objects/example.txt"
  content_type = "text/plain"
  bucket       = google_storage_bucket.static.id
}