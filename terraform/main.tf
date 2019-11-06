locals {
  // TODO: replace <YOUR_PROJECT>
  project = "<YOUR_PROJECT>"
  dataset = "sample_sink_target"

  bq_location = "US"
}

resource "google_logging_project_sink" "my_sample_sink" {
  name = "my-sample-sink"
  project = local.project
  destination = "bigquery.googleapis.com/projects/${local.project}/datasets/${local.dataset}"
  // TODO: replace <YOUR_PROJECT>
  filter = "resource.type=project AND resource.labels.project_id=<YOUR_PROJECT> AND labels.bigquery-export-type=sample"
  unique_writer_identity = true
}

resource "google_bigquery_dataset" "dataset" {
  dataset_id = local.dataset
  project = local.project
  location = local.bq_location
  description = "sample sink"
}

resource "google_project_iam_member" "bigquery_sink_member" {
  project = local.project
  role = "roles/bigquery.dataEditor"
  member = google_logging_project_sink.my_sample_sink.writer_identity
}