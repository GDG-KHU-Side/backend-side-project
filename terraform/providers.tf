provider "google" {
  credentials = file("service_account.json")
  project     = "your-project-id"
  region      = "us-central1"
}