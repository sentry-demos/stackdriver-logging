# sentry-demos/stackdriver-logging

This demo covers sending stackdriver logs to Sentry.

The cloud function consumes Stackdriver logs by subscribing to a PubSub topic.

# First Time Setup

## Stackdriver

Create a Stackdriver logging filter - https://cloud.google.com/logging/docs/view/basic-filters

You can also create more advanced filters - https://cloud.google.com/logging/docs/view/advanced-queries

Afterwards, create an export sink, using the `Create Export` option at the top of the `Logs Viewer`.

Select Pub/Sub as your `Sink Service` and select `Create new Cloud Pub/Sub Topic` to create a new topic.

https://cloud.google.com/logging/docs/export/configure_export_v2#dest-create

![Log Export View](./images/sink-export.png?raw=true)

## Cloud Function

Create a cloud function that is triggered from `Cloud Pub/Sub` and is pulling from the topic you created above.

Copy the code in `function.go` and paste it into the inline editor and make sure to insert your Sentry DSN.

```go
sentry.Init(sentry.ClientOptions{
    Dsn: "INSERT_DSN_HERE",
})
```

https://cloud.google.com/functions/docs/calling/pubsub

![Cloud Function Creation Screen](./images/cloudfunction.png?raw=true)

# Further customization

The current cloud function is just an example of what you can do. For further functionality, it is recommended that messages themselves be parsed for tags, similar to what [sentlog](https://github.com/getsentry/sentlog) does.
