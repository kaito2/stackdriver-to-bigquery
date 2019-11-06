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