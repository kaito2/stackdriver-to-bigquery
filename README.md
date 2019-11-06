# なに

Cloud Functions のログを Stackdriver logging から BigQuery へシンクするサンプル

## Deploy

※ 簡単化のために `Compute Engine default service account` を使っている。

```
export PROJECT=`gcloud config get-value core/project`
export PROJECT_NUMBER=`gcloud projects describe $PROJECT --format="value(projectNumber)"`
export SERVICE_ACCOUNT=${PROJECT_NUMBER}-compute@developer.gserviceaccount.com

gcloud functions deploy Hello --runtime go111 --trigger-http \
    --set-env-vars PROJECT=${PROJECT} \
    --service-account ${SERVICE_ACCOUNT}
```

## References

- [Package logging](https://godoc.org/cloud.google.com/go/logging)
- [Exporting with the Logs Viewer &nbsp;|&nbsp; Stackdriver Logging
      &nbsp;|&nbsp; Google Cloud](https://cloud.google.com/logging/docs/export/configure_export_v2)
- [terraform-google-modules/terraform-google-log-export](https://github.com/terraform-google-modules/terraform-google-log-export)
- [Google: google_bigquery_dataset - Terraform by HashiCorp](https://www.terraform.io/docs/providers/google/r/bigquery_dataset.html)