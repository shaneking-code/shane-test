terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
    random = {
      source = "hashicorp/random"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = "us-east1"
}

resource "random_id" "unique_id" {
  byte_length = 4
}

resource "google_storage_bucket" "static" {
  name                        = "shane-test-bucket-${random_id.unique_id.hex}"
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